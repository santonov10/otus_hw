package internalhttp

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct { // TODO
	server *http.Server
}

func NewServer() *Server {
	r := gin.Default()
	mainGroup := r.Group("/")

	f, err := os.Create("gin.log")
	if err != nil {
		gin.DefaultWriter = io.MultiWriter(f)
	}

	mainGroup.Use(gin.LoggerWithFormatter(LoggingMiddleware))
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
