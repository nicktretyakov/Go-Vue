package main

import (
	//"backend/models"
	//"encoding/json"
	"errors"
	//"log"
	"net/http"
	"strconv"
	//"time"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneBookingParticipant(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("id is", id)

	participant, err := app.models.DB.GetBookingParticipant(id)

	err = app.writeJSON(w, http.StatusOK, participant, "participant")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

}

func (app *application) getAllBookingParticipants(w http.ResponseWriter, r *http.Request) {
	participants, err := app.models.DB.AllBookingParticipant()
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

// func (app *application) getAllBookings(w http.ResponseWriter, r *http.Request) {
// 	bookings, err := app.models.DB.BookingsAll()
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}

// 	err = app.writeJSON(w, http.StatusOK, bookings, "bookings")
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}
// }

func (app *application) getAllParticipantsByBooking(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	bookingID, err := strconv.Atoi(params.ByName("booking_id"))

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	participants, err := app.models.DB.AllBookingParticipant(bookingID)
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

// type ParticipantPayload struct {
// 	ID            string `json:"id"`
// 	ParticipantName      string `json:"participant_name"`
// 	ParticipantEmail     string `json:"participant_email"`
// 	SecurityEmail string `json:"security_email"`
// 	Visible       bool   `json:"visible"`
// }

// func (app *application) editParticipant(w http.ResponseWriter, r *http.Request) {

// 	var payload ParticipantPayload

// 	err := json.NewDecoder(r.Body).Decode(&payload)
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}

// 	log.Println(payload.ParticipantName)

// 	var participant models.Participant

// 	if payload.ID != "0" {
// 		id, _ := strconv.Atoi(payload.ID)
// 		m, _ := app.models.DB.GetBookingParticipant(id)
// 		participant = *m
// 		participant.UpdatedAt = time.Now()
// 	}

// 	participant.ID, _ = strconv.Atoi(payload.ID)
// 	participant.ParticipantName = payload.ParticipantName
// 	participant.Space = payload.Space
// 	participant.SecurityEmail = payload.SecurityEmail
// 	participant.CreatedAt = time.Now()
// 	participant.UpdatedAt = time.Now()

// 	if participant.ID == 0 {
// 		err = app.models.DB.InsertParticipant(participant)
// 		if err != nil {
// 			app.errorJSON(w, err)
// 			return
// 		}
// 	} else {
// 		err = app.models.DB.UpdateParticipant(participant)
// 		if err != nil {
// 			app.errorJSON(w, err)
// 			return
// 		}
// 	}

// 	ok := jsonResp{
// 		OK: true,
// 	}

// 	err = app.writeJSON(w, http.StatusOK, ok, "response")
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}
// }

 func (app *application) deleteParticipantBookings(w http.ResponseWriter, r *http.Request) {
 	params := httprouter.ParamsFromContext((r.Context()))

 	id, err := strconv.Atoi(params.ByName("id"))
 	if err != nil {
 		app.errorJSON(w, err)
 		return
 	}

 	err = app.models.DB.DeleteParticipantBookings(id)
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
