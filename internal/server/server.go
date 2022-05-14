package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rewle/service-select-participants/internal/config"
	"github.com/rewle/service-select-participants/internal/participant"
	"go.uber.org/zap"
)

type Server struct {
	participantRepo *participant.ParticipantRepository

	log  *zap.SugaredLogger
	addr string
}

func (s *Server) Run() {
	r := mux.NewRouter()

	s.participants(r)
	// calendar TODO

	s.log.Fatal(http.ListenAndServe(s.addr, r))
}

func Init(log *zap.SugaredLogger, cfg *config.Config, p *participant.ParticipantRepository) *Server {
	return &Server{p, log, cfg.Addr}
}
