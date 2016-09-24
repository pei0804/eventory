package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/tikasan/eventory/server/db"
	"github.com/tikasan/eventory/server/define"
	"github.com/tikasan/eventory/server/model"

	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	db *sql.DB
}

func New() *Server {
	return &Server{}
}

func (s *Server) CheckAndInsert() <-chan []model.Event {
	now := time.Now()
	atdn := make([]model.Inserter, define.SERACH_SCOPE)
	connpass := make([]model.Inserter, define.SERACH_SCOPE)
	//doorKeeper := make([]model.Inserter, define.SERACH_SCOPE)
	allInserter := make([]model.Inserter, 0)

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
	var wg sync.WaitGroup

	go func() {
		for _, a := range allInserter {
			wg.Add(1)
			go func(a model.Inserter) {
				cli := model.NewInserter(a.Url, a.Api)
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

func (s *Server) Response() []model.EventJson {
	Event, err := model.EventAllNew(s.db)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	return Event
}

func (s *Server) Init(dbconf, env string) {
	cs, err := db.NewConfigsFromFile(dbconf)
	if err != nil {
		log.Fatalf("cannot open database configuration. exit. %s", err)
	}
	s.db, err = cs.Open(env)
	if err != nil {
		log.Fatalf("db initialization failed: %s", err)
	}

	_, err = os.Stat("./log")
	if err != nil {
		err := os.Mkdir("./log", 0750)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			log.Fatalf("log folder initialization failed: %s", err)
		}

	}

	_, err = os.Stat("./log/admin.log")
	if err != nil {
		_, err := os.Create("./log/admin.log")
		if err != nil {
			log.Fatalf("log admin.log initialization failed: %s", err)
		}
	}
}

func (s *Server) Run(port string) {
	http.HandleFunc("/api/pc/events", func(w http.ResponseWriter, r *http.Request) {
		event := s.Response()
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)
		if err := enc.Encode(event); err != nil {
			http.Error(w, "encoding failed", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/api/smt/events", func(w http.ResponseWriter, r *http.Request) {
		event := s.Response()
		//for i := range event {
		//	event[i].Desc = formater.RemoveTag(event[i].Desc)
		//}

		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)

		if err := enc.Encode(event); err != nil {
			http.Error(w, "encoding failed", http.StatusInternalServerError)
			return
		}
		fmt.Println("きた")
	})

	http.HandleFunc("/api/events/admin", func(w http.ResponseWriter, r *http.Request) {
		adminLog, err := os.OpenFile("./log/admin.log", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			return
		}
		defer adminLog.Close()

		now := time.Now()
		logger := log.New(adminLog, "[start]", log.LstdFlags)
		logger.Println(now)
		adminLog.Sync()

		receiver := s.CheckAndInsert()
		for {
			receive, ok := <-receiver
			if ok {
				model.Insert(s.db, receive)
			} else {
				break
			}

		}

		end := time.Now()
		logger = log.New(adminLog, "[end]", log.LstdFlags)
		logger.Println(end)
		adminLog.Sync()
		return
	})
	log.Fatal(http.ListenAndServe(port, nil))
}

func main() {
	var (
		port   = flag.String("port", ":8080", "port to bind")
		dbconf = flag.String("dbconf", "dbconfig.yml", "database configuration file.")
		env    = flag.String("env", "development", "application envirionment (production, development etc.)")
	)

	flag.Parse()
	s := New()

	// イニシャライザが走る
	s.Init(*dbconf, *env)
	s.Run(*port)
}
