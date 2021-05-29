package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hokkung/go-gateway/server/router"
)

type Server struct {
	engine *gin.Engine
}

func New() *Server {
	engine := gin.Default()
	router.Register(engine)

	return &Server{
		engine: engine,
	}
}

func (s *Server) Run() {
	s.engine.Run()
}
