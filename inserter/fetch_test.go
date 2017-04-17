package inserter

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo"
	"github.com/mjibson/goon"
	"github.com/tikasan/eventory-old/formater"
	"github.com/tikasan/eventory-old/model"
	"github.com/tikasan/eventory/define"
)

// TODO ネーミング変えるべきかも
type Fetch struct {
	DB *sql.DB
}

func (f *Fetch) EventFetch(c context.Context) error {

	err := dataStoreCheck(c)
	if err != nil {
	}

	receiver := communication(c)

	for {
		receive, ok := <-receiver
		if !ok {
			break
		}
		err := model.Insert(f.DB, receive)
		if err != nil {
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
	f.DB.Close()

	return
}

func communication(c echo.Context) <-chan []model.Event {
	now := time.Now()
	atdn := make([]Parser, define.SERACH_SCOPE)
	connpass := make([]Parser, define.SERACH_SCOPE)
	doorKeeper := make([]Parser, define.SERACH_SCOPE)
	allRequest := make([]Parser, 0)

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
		go func(p Parser) {
			cli := NewParser(p.Url, p.Api, p.Token, c)
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

func (f *Fetch) GetEvent(c echo.Context) error {

	decodePlaces, err := formater.DecodeUriCompontent(c.QueryParam("places"))
	if err != nil {
		decodePlaces = ""
	}
	places := strings.Split(decodePlaces, ",")

	err = dataStoreCheck(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("[err][datastore init] %s", err))
	}

	updatedAt := c.QueryParam("updated_at")
	layout := "2006-01-02 15:04:05"

	uut, err := time.Parse(layout, updatedAt)
	if err != nil {
		uut, _ = time.Parse(layout, "2000-01-01 00:00:00")
	}

	g := goon.NewGoon(c.Request())
	u := model.UpdateInfo{Id: define.PRODUCTION}
	err = g.Get(&u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("[err][datastore -> time] %s", err))
	}

	if uut.Unix() >= u.Datetime.Unix() {
		return c.JSON(http.StatusNotModified, fmt.Sprintf("lastUpdate %s", u.Datetime))
	}

	event, err := model.EventAllNew(f.DB, updatedAt, places)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	c.Response().Header().Set("Content-Type", "application/json")
	f.DB.Close()
	return c.JSON(http.StatusOK, event)
}

func dataStoreCheck(c context.Context) error {

	g := goon.NewGoon(c.Request())
	u := model.UpdateInfo{Id: define.PRODUCTION}
	err := g.Get(&u)
	if err != nil {
		u := model.UpdateInfo{Id: define.PRODUCTION, Datetime: time.Now()}
		_, err = g.Put(&u)
		if err != nil {
			return err
		}
	}
	return nil
}
