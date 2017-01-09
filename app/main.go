package server

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tikasan/eventory/api"
	"github.com/tikasan/eventory/db"

	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	db *sql.DB
}

func New() *Server {
	return &Server{}
}

func (s *Server) Setup(dbconf, env string) {

	cs, err := db.NewConfigsFromFile(dbconf)
	if err != nil {
		log.Fatalf("cannot open database configuration. exit. %s", err)
	}
	s.db, err = cs.Open(env)
	if err != nil {
		log.Fatalf("db initialization failed: %s", err)
	}
}

func (s *Server) Run() {

	api := &api.Inserter{DB: s.db}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/api/smt/events", api.GetEvent)
	e.GET("/api/events/admin", api.EventFetch)

	e.Pre(middleware.RemoveTrailingSlash())
	http.Handle("/", e)
}

func init() {
	var (
		dbconf = flag.String("dbconf", "dbconfig.yml", "database configuration file.")
		env    = flag.String("env", "development", "application envirionment (production, development etc.)")
	)

	flag.Parse()
	s := New()
	s.Setup(*dbconf, *env)
	s.Run()
}
