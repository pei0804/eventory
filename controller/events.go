package controller

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/tikasan/eventory/app"
	"github.com/tikasan/eventory/models"
)

// EventsController implements the events resource.
type EventsController struct {
	*goa.Controller
	db *gorm.DB
}

// NewEventsController creates a events controller.
func NewEventsController(service *goa.Service, db *gorm.DB) *EventsController {
	return &EventsController{
		Controller: service.NewController("EventsController"),
		db:         db,
	}
}

// List runs the list action.
func (c *EventsController) List(ctx *app.ListEventsContext) error {
	// EventsController_List: start_implement

	// Put your logic here
	eventDB := models.NewEventDB(c.db)
	events, err := eventDB.ListByQ(ctx.Context, ctx.Q, ctx.Sort, ctx.Page)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	// EventsController_List: end_implement
	return ctx.OK(events)
}

// ユーザーのキープ操作
func (c *EventsController) Keep(ctx *app.KeepEventsContext) error {
	// EventsController_Keep: start_implement

	// Put your logic here

	// EventsController_Keep: end_implement
	return nil
}
