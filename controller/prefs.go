package controller

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/tikasan/eventory/app"
)

// PrefsController implements the prefs resource.
type PrefsController struct {
	*goa.Controller
	db *gorm.DB
}

// NewPrefsController creates a prefs controller.
func NewPrefsController(service *goa.Service, db *gorm.DB) *PrefsController {
	return &PrefsController{
		Controller: service.NewController("PrefsController"),
		db:         db,
	}
}

// PrefFollow runs the pref follow action.
func (c *PrefsController) PrefFollow(ctx *app.PrefFollowPrefsContext) error {
	// PrefsController_PrefFollow: start_implement

	// Put your logic here

	// PrefsController_PrefFollow: end_implement
	return nil
}
