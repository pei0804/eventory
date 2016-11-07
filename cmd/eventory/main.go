package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/tikasan/eventory/db"

	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tikasan/eventory/api"
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

	g, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	logDir := filepath.Join(g, "log")
	_, err = os.Stat(logDir)
	if err != nil {
		err := os.Mkdir(logDir, 0750)
		if err != nil {
			log.Fatalf("log folder initialization failed: %s", err)
		}

	}

	checkLogPath := filepath.Join(logDir, "check.log")
	_, err = os.Stat(checkLogPath)
	if err != nil {
		_, err := os.Create(checkLogPath)
		if err != nil {
			log.Fatal("log check.log initialization failed: %s", err)
		}
	}
	return
}

func (s *Server) Run(port string) {
	http.HandleFunc("/api/smt/events", func(w http.ResponseWriter, r *http.Request) {
		api.Response(w, s.db)
		return
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

	s.Init(*dbconf, *env)
	s.Run(*port)
}
