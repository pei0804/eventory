package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"database/sql"

	"github.com/tikasan/eventory/server/define"
	"github.com/tikasan/eventory/server/model"
)

func Check(db *sql.DB) {

	g, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	checkLogPath := filepath.Join(g, "log", "check.log")
	_, err = os.Stat(checkLogPath)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	checkLog, err := os.OpenFile(checkLogPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return
	}
	defer checkLog.Close()

	now := time.Now()
	logger := log.New(checkLog, "[start]", log.LstdFlags)
	logger.Println(now)
	checkLog.Sync()

	receiver := Request()
	for {
		receive, ok := <-receiver
		if ok {
			model.Insert(db, receive)
		} else {
			break
		}

	}

	end := time.Now()
	logger = log.New(checkLog, "[end]", log.LstdFlags)
	logger.Println(end)
	checkLog.Sync()
}

func Request() <-chan []model.Event {
	now := time.Now()
	atdn := make([]model.Inserter, define.SERACH_SCOPE)
	connpass := make([]model.Inserter, define.SERACH_SCOPE)
	doorKeeper := make([]model.Inserter, define.SERACH_SCOPE)
	allInserter := make([]model.Inserter, 0)

	for i := 0; i < define.SERACH_SCOPE; i++ {
		ym := now.AddDate(0, i, 0).Format("200601")
		atdn[i].Url = fmt.Sprintf("https://api.atnd.org/events/?count=100&format=jsonp&callback=&ym=%s", ym)
		atdn[i].Api = define.ATDN

		connpass[i].Url = fmt.Sprintf("https://connpass.com/api/v1/event/?count=100&ym=%s", ym)
		connpass[i].Api = define.CONNPASS

		doorKeeper[i].Url = fmt.Sprintf("https://api.doorkeeper.jp/events?page=%d", i)
		doorKeeper[i].Api = define.DOORKEEPER
		doorKeeper[i].Token = "Bearer key"
	}

	allInserter = append(allInserter, atdn...)
	allInserter = append(allInserter, connpass...)
	allInserter = append(allInserter, doorKeeper...)
	allEvents := make(chan []model.Event, len(allInserter))
	var wg sync.WaitGroup

	go func() {
		for _, a := range allInserter {
			wg.Add(1)
			go func(a model.Inserter) {
				cli := model.NewInserter(a.Url, a.Api, a.Token)
				events, err := cli.Get()
				if err != nil {
					fmt.Fprint(os.Stderr, err)
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

func Response(w http.ResponseWriter, db *sql.DB) {
	event, err := model.EventAllNew(db)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	if err := enc.Encode(event); err != nil {
		http.Error(w, "encoding failed", http.StatusInternalServerError)
		return
	}
}
