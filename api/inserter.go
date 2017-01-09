package api

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"

	"database/sql"

	"github.com/labstack/echo"
	"github.com/tikasan/eventory/define"
	"github.com/tikasan/eventory/model"
)

// TODO ネーミング変えるべきかも
type Inserter struct {
	DB *sql.DB
}

func (i *Inserter) EventFetch(c echo.Context) error {

	if c.FormValue("token") != "rzo23y_fgRK1hnDDvMAH" {
		return c.JSON(http.StatusUnauthorized, "[err][auth check]")
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

	// TODO 本来は下のメソッドを先に実行すべき。+　全てチェックするべき、リリース後対応する。
	ctx := appengine.NewContext(c.Request())
	client := urlfetch.Client(ctx)

	_, err := client.Head(define.ATDN_URL)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("[err][atdn cant access]", err))
	}
	_, err = client.Head(define.CONNPASS_URL)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("[err][connpass cant access] %s", err))
	}
	_, err = client.Head(define.DOORKEEPER_URL)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("[err][doorkeeper cant access] %s", err))
	}

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
	event, err := model.EventAllNew(i.DB, updatedAt)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(http.StatusOK, event)
}
