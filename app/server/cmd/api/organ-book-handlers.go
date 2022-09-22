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


func (app *application) getOneBookingOrganizer(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("id is", id)

	organizer, err := app.models.DB.GetBookingOrganizer(id)

	err = app.writeJSON(w, http.StatusOK, organizer, "organizer")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

}

func (app *application) getAllBookingOrganizers(w http.ResponseWriter, r *http.Request) {
	organizers, err := app.models.DB.AllBookingOrganizer()
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

func (app *application) getAllOrganizersByBooking(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	bookingID, err := strconv.Atoi(params.ByName("booking_id"))

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	organizers, err := app.models.DB.AllBookingOrganizer(bookingID)
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

type OrganizerPayload struct {
	ID            string `json:"id"`
	OrganizerName string `json:"organizer_name"`
	OrganizerEmail string    `json:"organizer_email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	
}

func (app *application) editOrganizer(w http.ResponseWriter, r *http.Request) {

	var payload OrganizerPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	log.Println(payload.OrganizerName)

	var organizer models.Organizer

	if payload.ID != "0" {
		id, _ := strconv.Atoi(payload.ID)
		m, _ := app.models.DB.GetBookingOrganizer(id)
		organizer = *m
		organizer.UpdatedAt = time.Now()
	}

	organizer.ID, _ = strconv.Atoi(payload.ID)
	organizer.OrganizerName = payload.OrganizerName
	organizer.OrganizerEmail = payload.OrganizerEmail
	organizer.CreatedAt = time.Now()
	organizer.UpdatedAt = time.Now()

	if organizer.ID == 0 {
		err = app.models.DB.InsertOrganizer(organizer)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
	} else {
		err = app.models.DB.UpdateOrganizer(organizer)
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

func (app *application) deleteOrganizer(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext((r.Context()))

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.models.DB.DeleteOrganizer(id)
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
