package controller

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/tikasan/eventory/app"
)

// CronController implements the cron resource.
type CronController struct {
	*goa.Controller
	db *gorm.DB
}

// NewCronController creates a cron controller.
func NewCronController(service *goa.Service, db *gorm.DB) *CronController {
	return &CronController{
		Controller: service.NewController("CronController"),
		db:         db,
	}
}

// AppendGenre runs the append genre action.
func (c *CronController) AppendGenre(ctx *app.AppendGenreCronContext) error {
	// CronController_AppendGenre: start_implement

	// Put your logic here

	// CronController_AppendGenre: end_implement
	return nil
}

// FixUserFollow runs the fix user follow action.
func (c *CronController) FixUserFollow(ctx *app.FixUserFollowCronContext) error {
	// CronController_FixUserFollow: start_implement

	// Put your logic here

	// CronController_FixUserFollow: end_implement
	return nil
}

// NewEventFetch runs the new event fetch action.
func (c *CronController) NewEventFetch(ctx *app.NewEventFetchCronContext) error {
	// CronController_NewEventFetch: start_implement

	// Put your logic here

	// CronController_NewEventFetch: end_implement
	return nil
}
