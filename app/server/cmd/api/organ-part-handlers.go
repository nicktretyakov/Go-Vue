package main

import (
	"backend/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneOrganizerParticipant(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("id is", id)

	participant, err := app.models.DB.GetOrganizerParticipant(id)

	err = app.writeJSON(w, http.StatusOK, participant, "participant")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

}

func (app *application) getAllOrganizerParticipants(w http.ResponseWriter, r *http.Request) {
	participants, err := app.models.DB.AllOrganizerParticipant()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, participants, "participants")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllOrganizers(w http.ResponseWriter, r *http.Request) {
	organizers, err := app.models.DB.OrganizersAll()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, organizers, "organizers")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllParticipantsByOrganizer(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	organizerID, err := strconv.Atoi(params.ByName("organizer_id"))

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	participants, err := app.models.DB.AllOrganizerParticipant(organizerID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, participants, "participants")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

type ParticipantPayload struct {
	ID            string `json:"id"`
	ParticipantName string `json:"participant_name"`
	ParticipantEmail   string     `json:"participant_email"`
	ParticipantCar string `json:"participant_car"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (app *application) editParticipant(w http.ResponseWriter, r *http.Request) {

	var payload ParticipantPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	log.Println(payload.ParticipantName)

	var participant models.Participant

	if payload.ID != "0" {
		id, _ := strconv.Atoi(payload.ID)
		m, _ := app.models.DB.GetOrganizerParticipant(id)
		participant = *m
		participant.UpdatedAt = time.Now()
	}

	participant.ID, _ = strconv.Atoi(payload.ID)
	participant.ParticipantName = payload.ParticipantName
	participant.ParticipantEmail = payload.ParticipantEmail
	participant.ParticipantCar = payload.ParticipantCar
	participant.CreatedAt = time.Now()
	participant.UpdatedAt = time.Now()

	if participant.ID == 0 {
		err = app.models.DB.InsertParticipant(participant)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
	} else {
		err = app.models.DB.UpdateParticipant(participant)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
	}

	ok := jsonResp{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) deleteParticipantOrganizers(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext((r.Context()))

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.models.DB.DeleteParticipantOrganizers(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	ok := jsonResp{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}
