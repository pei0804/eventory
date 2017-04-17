// +build appengine

//go:generate goagen bootstrap -d github.com/tikasan/eventory/design

package server

import (
	"log"
	"net/http"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jinzhu/gorm"
	"github.com/tikasan/eventory/app"
	"github.com/tikasan/eventory/controller"
	"github.com/tikasan/eventory/database"
	"github.com/tikasan/eventory/models"
	"github.com/tikasan/eventory/utility"
)

func init() {
	// Create service
	service := goa.New("eventory")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	cs, err := database.NewConfigsFromFile("dbconfig.yml")
	if err != nil {
		log.Fatalf("cannot open database configuration. exit. %s", err)
	}
	dbcon, err := cs.Open("setting")
	if err != nil {
		log.Fatalf("database initialization failed: %s", err)
	}

	app.UseUserTokenMiddleware(service, NewAPIKeyMiddleware(dbcon))
	app.UseCronTokenMiddleware(service, NewCronAuthKeyMiddleware())

	// Mount "events" controller
	c := controller.NewEventsController(service, dbcon)
	app.MountEventsController(service, c)
	// Mount "genres" controller
	c2 := controller.NewGenresController(service, dbcon)
	app.MountGenresController(service, c2)
	// Mount "prefs" controller
	c3 := controller.NewPrefsController(service, dbcon)
	app.MountPrefsController(service, c3)
	// Mount "users" controller
	c4 := controller.NewUsersController(service, dbcon)
	app.MountUsersController(service, c4)
	// Mount "cron" controller
	c5 := controller.NewCronController(service, dbcon)
	app.MountCronController(service, c5)

	// Setup HTTP handler
	http.HandleFunc("/", service.Mux.ServeHTTP)
}

func NewAPIKeyMiddleware(db *gorm.DB) goa.Middleware {

	// Instantiate API Key security scheme details generated from design
	scheme := app.NewUserTokenSecurity()

	// Middleware
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// Retrieve and log header specified by scheme
			token := req.Header.Get(scheme.Name)
			// A real app would do something more interesting here
			if len(token) == 0 {
				goa.LogInfo(ctx, "failed api token auth")
				return errors.Unauthenticated("missing auth")
			}
			userTerminalDB := models.NewUserTerminalDB(db)
			userID, err := userTerminalDB.GetUserIDByToken(ctx, token)
			if err != nil {
				goa.LogInfo(ctx, "failed api token auth")
				return errors.Unauthenticated("missing auth")
			}
			utility.SetUserID(ctx, userID)
			utility.SetToken(ctx, token)
			// Proceed.
			goa.LogInfo(ctx, "auth", "apikey", "token", token)
			return h(ctx, rw, req)
		}

	}
}

func NewCronAuthKeyMiddleware() goa.Middleware {

	// Instantiate API Key security scheme details generated from design
	scheme := app.NewCronTokenSecurity()

	// Middleware
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// Retrieve and log header specified by scheme
			token := req.Header.Get(scheme.Name)
			// A real app would do something more interesting here
			if len(token) == 0 {
				goa.LogInfo(ctx, "failed api token auth")
				return errors.Unauthenticated("missing auth")
			}
			// Proceed.
			goa.LogInfo(ctx, "auth", "apikey", "token", token)
			return h(ctx, rw, req)
		}

	}
}
