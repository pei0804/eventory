package server

import (
	"database/sql"
	"flag"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tikasan/eventory/api"

	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	db   *sql.DB
	echo *echo.Echo
}

func New() *Server {
	return &Server{}
}

func (s *Server) Run() {
	api := &api.Inserter{DB: s.db}

	s.echo = echo.New()

	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.CORS())

	s.echo.GET("/api/smt/events", api.GetEvent)
	s.echo.GET("/api/events/admin", api.EventFetch)

	s.echo.Pre(middleware.RemoveTrailingSlash())
	http.Handle("/", s.echo)
}

func init() {
	flag.Parse()
	s := New()
	s.Run()
}
