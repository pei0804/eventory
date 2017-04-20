package controller

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/tikasan/eventory/app"
	"github.com/tikasan/eventory/models"
	"github.com/tikasan/eventory/utility"
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

// フォロー、アンフォロー操作
func (c *PrefsController) Follow(ctx *app.FollowPrefsContext) error {
	// PrefsController_Follow: start_implement

	// Put your logic here
	ufg := &models.UserFollowPref{}
	ufg.PrefID = ctx.PrefID
	userID, err := utility.GetUserID(ctx.Context)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	ufg.UserID = userID
	ufgDB := models.NewUserFollowPrefDB(c.db)
	// HTTPメソッドで判定する
	if "PUT" == ctx.Request.Method {
		ufgDB.UserFollowPref(ctx.Context, ufg)
	}
	if "DELETE" == ctx.Request.Method {
		ufgDB.UserUnfollowPref(ctx.Context, ufg)
	}
	// PrefsController_Follow: end_implement
	return nil
}
