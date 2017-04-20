package controller

import (
	"fmt"
	"os"
	"time"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/tikasan/eventory/app"
	"github.com/tikasan/eventory/define"
	"github.com/tikasan/eventory/inserter"
	"github.com/tikasan/eventory/models"
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

// ユーザーのキープ操作を確定させる。範囲は一日以上前の操作
func (c *CronController) FixUserKeep(ctx *app.FixUserKeepCronContext) error {
	// CronController_FixUserKeep: start_implement

	// Put your logic here
	ufe := models.NewUserKeepStatusDB(c.db)
	err := ufe.FixUserKeep(ctx.Context)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	// CronController_FixUserKeep: end_implement
	return nil
}

// 新しいイベント情報を取得する
func (c *CronController) NewEventFetch(ctx *app.NewEventFetchCronContext) error {
	// CronController_NewEventFetch: start_implement

	// Put your logic here
	now := time.Now()
	atdn := make([]inserter.Parser, define.SERACH_SCOPE)
	connpass := make([]inserter.Parser, define.SERACH_SCOPE)
	doorKeeper := make([]inserter.Parser, define.SERACH_SCOPE)

	for i := 0; i < define.SERACH_SCOPE; i++ {
		ym := now.AddDate(0, i, 0).Format("200601")
		atdn[i].URL = fmt.Sprintf("%s&ym=%s", define.ATDN_URL, ym)
		atdn[i].APIType = define.ATDN
		connpass[i].URL = fmt.Sprintf("%s&ym=%s", define.CONNPASS_URL, ym)
		connpass[i].APIType = define.CONNPASS
		doorKeeper[i].URL = fmt.Sprintf("%s?page=%d", define.DOORKEEPER_URL, i)
		doorKeeper[i].APIType = define.DOORKEEPER
		doorKeeper[i].Token = ""
	}
	allParser := make([]inserter.Parser, 0)
	allParser = append(allParser, atdn...)
	allParser = append(allParser, connpass...)
	allParser = append(allParser, doorKeeper...)

	// TODO: API提供元に負荷がかかるので、順次処理にしている。
	for _, p := range allParser {
		cli := inserter.NewParser(p.URL, p.APIType, p.Token, ctx.Request)
		es, err := cli.ConvertingToJson()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		for _, e := range es {
			events := &models.Event{}
			events.ID = e.ID
			events.APIType = e.APIType
			events.Address = e.Title
			events.Accept = e.Accept
			events.Identifier = e.Identifier
			events.DataHash = e.DataHash
			events.Description = e.Description
			events.Limits = e.Limits
			events.Pref = e.Pref
			events.URL = e.URL
			events.Wait = e.Wait
			events.StartAt = e.StartAt
			events.EndAt = e.EndAt
			eventsDB := models.NewEventDB(c.db)
			err := eventsDB.Add(ctx.Context, events)
			if err != nil {
				return fmt.Errorf("%v", err)
			}
		}
	}
	// CronController_NewEventFetch: end_implement
	return nil
}
