package controller

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/tikasan/eventory/app"
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

	// GenresController_Create: end_implement
	res := &app.Genre{}
	return ctx.OK(res)
}

// FollowGenre runs the follow genre action.
func (c *GenresController) FollowGenre(ctx *app.FollowGenreGenresContext) error {
	// GenresController_FollowGenre: start_implement

	// Put your logic here

	// GenresController_FollowGenre: end_implement
	return nil
}

// List runs the list action.
func (c *GenresController) List(ctx *app.ListGenresContext) error {
	// GenresController_List: start_implement

	// Put your logic here

	// GenresController_List: end_implement
	res := app.GenreCollection{}
	return ctx.OK(res)
}
