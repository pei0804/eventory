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
type Inserter struct {
	DB *sql.DB
}

func (i *Inserter) EventFetch(c echo.Context) error {

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

	_, err = http.Head(define.ATDN_URL)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "atdn cant api access")
	}
	_, err = http.Head(define.CONNPASS_URL)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "connpass cant api access")
	}
	_, err = http.Head(define.DOORKEEPER_URL)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "doorkeeper cant api access")
	}

	receiver := communication()

	for {
		receive, ok := <-receiver
		if !ok {
			break
		} else {
			err := model.Insert(i.DB, receive)
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

func communication() <-chan []model.Event {
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

	go func() {
		for _, r := range allRequest {
			wg.Add(1)
			go func(r Request) {
				cli := NewRequest(r.Url, r.Api, r.Token)
				events, err := cli.Get()
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
	}()
	return allEvents
}

func (i *Inserter) GetEvent(c echo.Context) error {
	event, err := model.EventAllNew(i.DB)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(http.StatusOK, event)
}
