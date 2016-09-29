package main

import (
	"database/sql"
	"encoding/csv"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/tikasan/eventory/server/db"

	"path/filepath"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tikasan/eventory/server/api"
	sr "github.com/tuvistavie/securerandom"
)

type Server struct {
	db       *sql.DB
	filePath string
	token    string
}

func New() *Server {
	return &Server{}
}

func (s *Server) Init(dbconf, env string) {

	// DB
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

	// API auth header
	authPath := filepath.Join(g, "auth.csv")
	_, err = os.Stat(authPath)
	if err != nil {
		_, err = os.Create(authPath)
		if err != nil {
			log.Fatal("auth.txt initialization failed: %s", err)
		}
	}

	authFile, err := os.Create(authPath)
	b, err := sr.Base64(20, true)
	if err != nil {
		log.Fatal("auth token create initialization failed: %s", err)
	}
	authFile.WriteString(b)

	s.token = b
	s.filePath = authPath

	// log
	logDir := filepath.Join(g, "log")
	_, err = os.Stat(logDir)
	if err != nil {
		err := os.Mkdir(logDir, 0750)
		if err != nil {
			log.Fatalf("log folder initialization failed: %s", err)
		}
	}

	checkLogPath := filepath.Join(logDir, "check.log")
	fmt.Println(checkLogPath)
	_, err = os.Stat(checkLogPath)
	if err != nil {
		_, err := os.Create(checkLogPath)
		if err != nil {
			log.Fatal("log check.log initialization failed: %s", err)
		}
	}
	return
}

func (s *Server) RemakeToken() {

	authFile, err := os.Create(s.filePath)
	b, err := sr.Base64(20, true)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	authFile.WriteString(b)

	// initが通っているので、プロセスを止めるエラー処理をしない。
	fp, err := os.Open(s.filePath)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	reader := csv.NewReader(fp)
	record, err := reader.Read()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	s.token = record[0]
	fmt.Printf("curl -H 'Auth:%s' http://localhost:8080/api/events/admin", s.token)
}

func (s *Server) Run(port string) {

	// アクセスに3回連続失敗したら、トークンを再発行する
	s.RemakeToken()
	faileCount := 3
	http.HandleFunc("/api/events/admin", func(w http.ResponseWriter, r *http.Request) {

		if s.token != r.Header.Get("Auth") {
			faileCount--
			if faileCount <= 0 {
				s.RemakeToken()
			}
		} else {
			api.Check(s.db)
		}
		return
	})

	http.HandleFunc("/api/smt/events", func(w http.ResponseWriter, r *http.Request) {
		api.Response(w, s.db)
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
