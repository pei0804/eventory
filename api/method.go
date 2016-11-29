package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"database/sql"

	"github.com/labstack/echo"
	"github.com/tikasan/eventory/define"
	"github.com/tikasan/eventory/model"
)

// TODO ネーミング変えるべきかも
type Method struct {
	DB *sql.DB
}

func (m *Method) Check(c echo.Context) error {

	g, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	checkLogPath := filepath.Join(g, "log", "check.log")
	_, err = os.Stat(checkLogPath)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return c.JSON(http.StatusInternalServerError, "log/check.log not found")
	}

	checkLog, err := os.OpenFile(checkLogPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "log/check.log cant open")
	}
	defer checkLog.Close()

	now := time.Now()
	logger := log.New(checkLog, "[start]", log.LstdFlags)
	logger.Println(now)
	checkLog.Sync()

	receiver := Request()
	for {
		receive, ok := <-receiver
		if !ok {
			break
		} else {
			err := model.Insert(m.DB, receive)
			if err != nil {
				end := time.Now()
				logger = log.New(checkLog, "[database error]", log.LstdFlags)
				logger.Println(end)
				checkLog.Sync()
				return c.JSON(http.StatusInternalServerError, "Database Insert Error")
			}
		}

	}

	end := time.Now()
	logger = log.New(checkLog, "[end]", log.LstdFlags)
	logger.Println(end)
	checkLog.Sync()
	return c.JSON(http.StatusOK, "OK")
}

func Request() <-chan []model.Event {
	now := time.Now()
	atdn := make([]Inserter, define.SERACH_SCOPE)
	connpass := make([]Inserter, define.SERACH_SCOPE)
	doorKeeper := make([]Inserter, define.SERACH_SCOPE)
	allInserter := make([]Inserter, 0)

	for i := 0; i < define.SERACH_SCOPE; i++ {
		ym := now.AddDate(0, i, 0).Format("200601")
		atdn[i].Url = fmt.Sprintf("https://api.atnd.org/events/?count=100&format=jsonp&callback=&ym=%s", ym)
		atdn[i].Api = define.ATDN

		connpass[i].Url = fmt.Sprintf("https://connpass.com/api/v1/event/?count=100&ym=%s", ym)
		connpass[i].Api = define.CONNPASS

		doorKeeper[i].Url = fmt.Sprintf("https://api.doorkeeper.jp/events?page=%d", i)
		doorKeeper[i].Api = define.DOORKEEPER
		doorKeeper[i].Token = "Bearer "
	}

	allInserter = append(allInserter, atdn...)
	allInserter = append(allInserter, connpass...)
	allInserter = append(allInserter, doorKeeper...)
	allEvents := make(chan []model.Event, len(allInserter))
	var wg sync.WaitGroup

	go func() {
		for _, a := range allInserter {
			wg.Add(1)
			go func(a Inserter) {
				cli := NewInserter(a.Url, a.Api, a.Token)
				events, err := cli.Get()
				if err != nil {
					fmt.Fprint(os.Stderr, err)
					wg.Done()
				}
				allEvents <- events
				wg.Done()
			}(a)
		}
		wg.Wait()
		close(allEvents)
	}()
	return allEvents
}

func (m *Method) Response(c echo.Context) error {
	event, err := model.EventAllNew(m.DB)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(http.StatusOK, event)
}
