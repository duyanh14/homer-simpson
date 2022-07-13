package server

import (
	"context"
	"fmt"
	"net/http"

	"golang-course/pkg/db"
	"golang-course/service/repository"

	"github.com/gin-gonic/gin"
)

//Server ...
type Server struct {
	httpServer *http.Server
	router     *gin.Engine
}

// NewServer construct server
func NewServer() (*Server, error) {

	router := gin.New()

	s := &Server{
		router: router,
	}

	return s, nil
}

// Init server
func (s *Server) Init() {
	ctx := context.Background()
	database, err := db.InitMysqlDB("sql6.freemysqlhosting.net", 3306, "sql6501022", "szkJpYD8aq", "sql6501022")
	if err != nil {
		panic(err)
	}
	repo := repository.NewRepo(database)
	domains := s.initDomains(ctx, repo)
	s.initRouters(domains)
}

//ListenHTTP ...
func (s *Server) ListenHTTP() error {
	address := fmt.Sprintf(":%s", "8010")

	s.httpServer = &http.Server{
		Handler: s.router,
		Addr:    address,
	}

	return s.httpServer.ListenAndServe()
}
