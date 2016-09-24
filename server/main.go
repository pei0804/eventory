package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tikasan/eventory/server/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tikasan/eventory/server/api"
)

type Server struct {
	db *sql.DB
}

func New() *Server {
	return &Server{}
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
	http.HandleFunc("/api/smt/events", func(w http.ResponseWriter, r *http.Request) {
		api.Response(w, s.db)
	})

	http.HandleFunc("/api/events/admin", func(w http.ResponseWriter, r *http.Request) {
		api.Check(s.db)
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
