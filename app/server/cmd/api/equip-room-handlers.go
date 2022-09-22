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

func (app *application) getOneRoomEquipment(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("id is", id)

	equipment, err := app.models.DB.GetRoomEquipment(id)

	err = app.writeJSON(w, http.StatusOK, equipment, "equipment")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

}

func (app *application) getAllRoomEquipments(w http.ResponseWriter, r *http.Request) {
	equipments, err := app.models.DB.AllRoomEquipment()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, equipments, "equipments")
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

func (app *application) getAllEquipmentsByRoom(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	roomID, err := strconv.Atoi(params.ByName("room_id"))

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	equipments, err := app.models.DB.AllRoomEquipment(roomID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, equipments, "equipments")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

type EquipmentPayload struct {
	ID            string `json:"id"`
	EquipmentName string `json:"city"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

func (app *application) editEquipment(w http.ResponseWriter, r *http.Request) {

	var payload EquipmentPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	log.Println(payload.EquipmentName)

	var equipment models.Equipment

	if payload.ID != "0" {
		id, _ := strconv.Atoi(payload.ID)
		m, _ := app.models.DB.GetRoomEquipment(id)
		equipment = *m
		equipment.UpdatedAt = time.Now()
	}

	equipment.ID, _ = strconv.Atoi(payload.ID)
	equipment.EquipmentName = payload.EquipmentName
	equipment.CreatedAt = time.Now()
	equipment.UpdatedAt = time.Now()

	if equipment.ID == 0 {
		err = app.models.DB.InsertEquipment(equipment)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
	} else {
		err = app.models.DB.UpdateEquipment(equipment)
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

func (app *application) deleteEquipment(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext((r.Context()))

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.models.DB.DeleteEquipment(id)
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
