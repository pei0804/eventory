package api

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"database/sql"

	"github.com/labstack/echo"
	"github.com/mjibson/goon"
	"github.com/tikasan/eventory/define"
	"github.com/tikasan/eventory/model"
)

// TODO ネーミング変えるべきかも
type Inserter struct {
	DB *sql.DB
}

func (i *Inserter) EventFetch(c echo.Context) error {

	if c.Request().Header.Get("X-Appengine-Cron")[0] == 0 {
		return c.JSON(http.StatusUnauthorized, fmt.Sprintf("[err][AuthError]"))
	}

	receiver := communication(c)

	for {
		receive, ok := <-receiver
		if !ok {
			break
		}
		err := model.Insert(i.DB, receive)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("[err][database insert] %s", err))
		}

	}

	//// TODO 本来は下のメソッドを先に実行すべき。+　全てチェックするべき、リリース後対応する。
	//ctx := appengine.NewContext(c.Request())
	//client := urlfetch.Client(ctx)
	//
	//_, err := client.Head(define.ATDN_URL)
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, fmt.Sprintf("[err][atdn cant access]", err))
	//}
	//_, err = client.Head(define.CONNPASS_URL)
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, fmt.Sprintf("[err][connpass cant access] %s", err))
	//}
	//_, err = client.Head(define.DOORKEEPER_URL)
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, fmt.Sprintf("[err][doorkeeper cant access] %s", err))
	//}

	g := goon.NewGoon(c.Request())
	u := model.UpdateInfo{Id: define.PRODUCTION, Datetime: time.Now()}
	g.Put(&u)

	return c.JSON(http.StatusOK, "OK")
}

func communication(c echo.Context) <-chan []model.Event {
	now := time.Now()
	atdn := make([]Request, define.SERACH_SCOPE)
	connpass := make([]Request, define.SERACH_SCOPE)
	doorKeeper := make([]Request, define.SERACH_SCOPE)
	allRequest := make([]Request, 0)

	for i := 0; i < define.SERACH_SCOPE; i++ {
		ym := now.AddDate(0, i, 0).Format("200601")
		atdn[i].Url = fmt.Sprintf("%s&ym=%s", define.ATDN_URL, ym)
		atdn[i].Api = define.ATDN

		connpass[i].Url = fmt.Sprintf("%s&ym=%s", define.CONNPASS_URL, ym)
		connpass[i].Api = define.CONNPASS

		doorKeeper[i].Url = fmt.Sprintf("%s?page=%d", define.DOORKEEPER_URL, i)
		doorKeeper[i].Api = define.DOORKEEPER
		doorKeeper[i].Token = ""
	}

	allRequest = append(allRequest, atdn...)
	allRequest = append(allRequest, connpass...)
	allRequest = append(allRequest, doorKeeper...)
	allEvents := make(chan []model.Event, len(allRequest))
	var wg sync.WaitGroup

	for _, r := range allRequest {
		wg.Add(1)
		go func(r Request) {
			cli := NewRequest(r.Url, r.Api, r.Token, c)
			events, err := cli.convertingToJson()
			if err != nil {
				fmt.Fprint(os.Stderr, err)
				wg.Done()
			}
			allEvents <- events
			wg.Done()
		}(r)
	}
	wg.Wait()
	close(allEvents)
	return allEvents
}

func (i *Inserter) GetEvent(c echo.Context) error {

	updatedAt := c.QueryParam("updated_at")
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, updatedAt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("[err][ios -> updateAt] %s", err))
	}
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	updateTime := t.In(jst)

	g := goon.NewGoon(c.Request())
	u := model.UpdateInfo{Id: define.PRODUCTION}
	err = g.Get(&u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("[err][datastore -> time] %s", err))
	}

	if !u.Datetime.After(updateTime) {
		return c.JSON(http.StatusNotModified, fmt.Sprintf("lastUpdate %s", u.Datetime))
	}

	event, err := model.EventAllNew(i.DB, updatedAt)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(http.StatusOK, event)
}
