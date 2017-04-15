package controller

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/tikasan/eventory/app"
	"github.com/tikasan/eventory/models"
	"github.com/tikasan/eventory/utility"
)

// GenresController implements the genres resource.
type GenresController struct {
	*goa.Controller
	db *gorm.DB
}

// NewGenresController creates a genres controller.
func NewGenresController(service *goa.Service, db *gorm.DB) *GenresController {
	return &GenresController{
		Controller: service.NewController("GenresController"),
		db:         db,
	}
}

// Create runs the create action.
func (c *GenresController) Create(ctx *app.CreateGenresContext) error {
	// GenresController_Create: start_implement

	// Put your logic here
	genre := &models.Genre{}
	genre.Name = ctx.Name
	genre.Keyword = ctx.Name
	genreDB := models.NewGenreDB(c.db)
	ID, err := genreDB.AddGetInsertID(ctx, genre)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	// 作成したGenreIDを返す
	// GenresController_Create: end_implement
	res := &app.GenreTiny{&ID}
	return ctx.OKTiny(res)
}

// Follow runs the follow action.
func (c *GenresController) Follow(ctx *app.FollowGenresContext) error {
	// GenresController_Follow: start_implement

	// Put your logic here
	ufg := &models.UserFollowGenre{}
	ufg.GenreID = ctx.GenreID
	userID, err := utility.GetUserID(ctx.Context)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	ufg.UserID = userID
	ufgDB := models.NewUserFollowGenreDB(c.db)
	if "PUT" == ctx.Request.Method {
		ufgDB.UserFollowGenre(ctx.Context, ufg)
	}

	if "DELETE" == ctx.Request.Method {
		ufgDB.UserUnfollowGenre(ctx.Context, ufg)
	}
	// GenresController_Follow: end_implement
	return nil
}

// List runs the list action.
func (c *GenresController) List(ctx *app.ListGenresContext) error {
	// GenresController_List: start_implement

	// Put your logic here
	genreDB := models.NewGenreDB(c.db)
	genre, err := genreDB.ListByKeyword(ctx.Context, ctx.Q, ctx.Sort, ctx.Page)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	// GenresController_List: end_implement
	return ctx.OK(genre)
}
