package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tikasan/eventory/server/db"
	"github.com/tikasan/eventory/server/define"
	"github.com/tikasan/eventory/server/formater"
	"github.com/tikasan/eventory/server/model"

	"time"

	_ "github.com/go-sql-driver/mysql"
)

func eventCheckAndInsert() {

	adminLog, err := os.OpenFile("./log/admin.log", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return
	}
	defer adminLog.Close()

	atdn := make([]model.Inserter, define.SERACH_SCOPE)
	connpass := make([]model.Inserter, define.SERACH_SCOPE)
	//doorKeeper := make([]model.Inserter, define.SERACH_SCOPE)
	allInserter := make([]model.Inserter, 0)

	now := time.Now()
	logger := log.New(adminLog, "[start]", log.LstdFlags)
	logger.Println(now)
	adminLog.Sync()

	for i := 0; i < define.SERACH_SCOPE; i++ {
		ym := now.AddDate(0, i, 0).Format("200601")
		atdn[i].Url = fmt.Sprintf("https://api.atnd.org/events/?count=100&format=jsonp&callback=&ym=%s", ym)
		atdn[i].Api = define.ATDN

		connpass[i].Url = fmt.Sprintf("https://connpass.com/api/v1/event/?count=100&ym=%s", ym)
		connpass[i].Api = define.CONNPASS

		//doorKeeper[i].Url = fmt.Sprintf("https://api.doorkeeper.jp/events?page=%d", i)
		//doorKeeper[i].Api = define.DOORKEEPER
	}

	allInserter = append(allInserter, atdn...)
	allInserter = append(allInserter, connpass...)
	//allInserter = append(allInserter, doorKeeper...)
	allEvents := make(chan []model.Event, len(allInserter))
	for _, a := range allInserter {
		go func(a model.Inserter) {
			cli := model.NewInserter(a.Url, a.Api)
			events, err := cli.Get()
			if err != nil {
				fmt.Fprint(os.Stderr, err)
			}
			allEvents <- events
		}(a)
	}
	sql := db.ConDB()
	defer sql.Close()
	for ev := range allEvents {
		model.Insert(sql, ev)
	}

	end := time.Now()
	logger = log.New(adminLog, "[end]", log.LstdFlags)
	logger.Println(end)
	adminLog.Sync()
}

func eventResponse() []model.EventJson {

	sql := db.ConDB()
	defer sql.Close()
	Event, err := model.EventAllNew(sql)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	return Event
}

func init() {

	// logファイルの存在　なければ作成
	_, err := os.Stat("./log")
	if err != nil {
		err := os.Mkdir("./log", 0750)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}

	}

	_, err = os.Stat("./log/admin.log")
	if err != nil {
		_, err := os.Create("./log/admin.log")
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}
	}
}

func main() {

	http.HandleFunc("/api/pc/events", func(w http.ResponseWriter, r *http.Request) {

		event := eventResponse()
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)
		if err := enc.Encode(event); err != nil {
			http.Error(w, "encoding failed", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/api/smt/events", func(w http.ResponseWriter, r *http.Request) {

		event := eventResponse()
		for i := range event {
			event[i].Desc = formater.RemoveTag(event[i].Desc)
		}

		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)

		if err := enc.Encode(event); err != nil {
			http.Error(w, "encoding failed", http.StatusInternalServerError)
			return
		}
		fmt.Println("きた")
	})

	http.HandleFunc("/api/events/admin", func(w http.ResponseWriter, r *http.Request) {
		eventCheckAndInsert()
		return
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
