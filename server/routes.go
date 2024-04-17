package server

import (
	"net/http"

	"github.com/ZafirProjects/QuodOrbisChallenge/handler"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	welcomeHandler := handler.WelcomeHandler{}
	mux.HandleFunc("GET /", welcomeHandler.HandleWelcomeRequest)

	generateHandler := handler.GenerateHandler{}
	mux.HandleFunc("POST /generate", generateHandler.RenderGeneratedData)
	mux.HandleFunc("POST /generate/new", generateHandler.RenderNewData)

	testHandler := handler.TestHandler{}
	mux.HandleFunc("GET /test", testHandler.HandleTestShow)

	return mux
}
