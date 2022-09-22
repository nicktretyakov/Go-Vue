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

type jsonResp struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (app *application) getOneRoom(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("id is", id)

	room, err := app.models.DB.GetBookingRoom(id)

	err = app.writeJSON(w, http.StatusOK, room, "room")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

}

func (app *application) getAllRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := app.models.DB.AllBookingRoom()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, rooms, "rooms")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := app.models.DB.BookingsAll()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, bookings, "bookings")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllRoomsByBooking(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	bookingID, err := strconv.Atoi(params.ByName("booking_id"))

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	rooms, err := app.models.DB.AllBookingRoom(bookingID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, rooms, "rooms")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

type RoomPayload struct {
	ID            string `json:"id"`
	RoomName      string `json:"room_name"`
	Space         int    `json:"space"`
	SecurityEmail string `json:"security_email"`
	Visible       bool   `json:"visible"`
}

func (app *application) editRoom(w http.ResponseWriter, r *http.Request) {

	var payload RoomPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	log.Println(payload.RoomName)

	var room models.Room

	if payload.ID != "0" {
		id, _ := strconv.Atoi(payload.ID)
		m, _ := app.models.DB.GetBookingRoom(id)
		room = *m
		room.UpdatedAt = time.Now()
	}

	room.ID, _ = strconv.Atoi(payload.ID)
	room.RoomName = payload.RoomName
	room.Space = payload.Space
	room.SecurityEmail = payload.SecurityEmail
	room.CreatedAt = time.Now()
	room.UpdatedAt = time.Now()

	if room.ID == 0 {
		err = app.models.DB.InsertRoom(room)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
	} else {
		err = app.models.DB.UpdateRoom(room)
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

func (app *application) deleteRoom(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext((r.Context()))

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.models.DB.DeleteRoom(id)
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
