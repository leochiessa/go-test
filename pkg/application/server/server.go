package server

import (
	"go-test/pkg/presentation/routers"
	"log"
	"net/http"
)

type Server struct {
	HTTP_PORT string
	AppRouter *routers.AppRouter
}

func (s *Server) Run() {
	log.Println("Listening on Port:", s.HTTP_PORT)
	err := http.ListenAndServe(":"+s.HTTP_PORT, s.AppRouter.Router)
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}
