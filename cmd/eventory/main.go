package main

import (
	"database/sql"
	"flag"
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tikasan/eventory/api"
	"github.com/tikasan/eventory/db"

	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
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

	api := &api.Method{DB: s.db}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/api/smt/events", api.Response)
	e.GET("/api/events/admin", api.Check)

	if err := e.Start(port); err != nil {
		e.Logger.Fatal(err.Error())
	}
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
