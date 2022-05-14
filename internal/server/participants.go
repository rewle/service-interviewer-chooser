package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rewle/service-select-participants/internal/participant"
	"github.com/rewle/service-select-participants/internal/utils"
)

func (s *Server) participants(r *mux.Router) {
	r.HandleFunc("/participant", func(w http.ResponseWriter, r *http.Request) {
		list, err := s.participantRepo.List()
		if err != nil {
			utils.BadGateway(w)
			return
		}
		status := http.StatusOK
		if len(*list) == 0 {
			status = http.StatusNotFound
		}
		utils.JSONResponse(w, status, list)
	}).Methods(http.MethodGet)

	r.HandleFunc("/participant", func(w http.ResponseWriter, r *http.Request) {
		data := participant.ParticipantCreateRequest{}

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&data)
		if err != nil {
			utils.BadGateway(w) // TODO bad response
			return
		}

		resp, err := s.participantRepo.Create(&data)
		if err != nil {
			utils.BadGateway(w)
			return
		}
		utils.JSONResponse(w, http.StatusOK, resp)
	}).Methods(http.MethodPost)

	r.HandleFunc("/participant", func(w http.ResponseWriter, r *http.Request) {
		data := participant.ParticipantRemoveRequest{}

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&data)
		if err != nil {
			utils.BadGateway(w)
			return
		}

		resp, err := s.participantRepo.Remove(&data)
		if err != nil {
			utils.BadGateway(w)
			return
		}
		utils.JSONResponse(w, http.StatusOK, resp)
	}).Methods(http.MethodDelete)
}
