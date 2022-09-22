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

func (app *application) getOneRoomAddress(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("id is", id)

	address, err := app.models.DB.GetRoomAddress(id)

	err = app.writeJSON(w, http.StatusOK, address, "address")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

}

func (app *application) getAllRoomAddresses(w http.ResponseWriter, r *http.Request) {
	addresses, err := app.models.DB.AllRoomAddress()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, addresses, "addresses")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

// func (app *application) getAllRooms(w http.ResponseWriter, r *http.Request) {
// 	rooms, err := app.models.DB.RoomsAll()
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}

// 	err = app.writeJSON(w, http.StatusOK, rooms, "rooms")
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}
// }

func (app *application) getAllAddressesByRoom(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	roomID, err := strconv.Atoi(params.ByName("room_id"))

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	addresses, err := app.models.DB.AllRoomAddress(roomID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, addresses, "addresses")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

type AddressPayload struct {
	ID            string `json:"id"`
	City          string `json:"city"`
	Street        string `json:"street"`
	Building      string `json:"building"`
	Floor         string `json:"floor"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

func (app *application) editAddress(w http.ResponseWriter, r *http.Request) {

	var payload AddressPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	log.Println(payload.City)

	var address models.Address

	if payload.ID != "0" {
		id, _ := strconv.Atoi(payload.ID)
		m, _ := app.models.DB.GetRoomAddress(id)
		address = *m
		address.UpdatedAt = time.Now()
	}

	address.ID, _ = strconv.Atoi(payload.ID)
	address.City = payload.City
	address.Street = payload.Street
	address.Building = payload.Building
	address.Floor = payload.Floor
	address.CreatedAt = time.Now()
	address.UpdatedAt = time.Now()

	if address.ID == 0 {
		err = app.models.DB.InsertAddress(address)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
	} else {
		err = app.models.DB.UpdateAddress(address)
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

func (app *application) deleteAddress(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext((r.Context()))

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.models.DB.DeleteAddress(id)
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
