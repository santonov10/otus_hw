package internalhttp

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/logger"

	"github.com/gin-gonic/gin"
)

type Server struct { // TODO
	server *http.Server
}

func NewServer() *Server {
	r := gin.Default()
	mainGroup := r.Group("/")

	mainGroup.Use(LoggingMiddleware(logger.Get()))
	{
		mainGroup.GET("/", mainPage)
	}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{
		server: s,
	}
}

func (s *Server) Start(ctx context.Context) {
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()
	go func() {
		<-ctx.Done()
		s.Stop(ctx)
	}()
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.server.Shutdown(ctx)
	fmt.Println("закрываем сервер")
	return err
}

func mainPage(c *gin.Context) {
	c.String(200, "Hello")
}
