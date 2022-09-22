package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	//"github.com/justinas/alice"
)

func (app *application) wrap(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}


func (app *application) routes() http.Handler {

	router := httprouter.New()

	//secure := alice.New(app.checkToken)

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodPost, "/v1/signin", app.Signin)
    router.HandlerFunc(http.MethodPost, "/v1/signin", app.SendMail)

	router.HandlerFunc(http.MethodGet, "/v1/room/:id", app.getOneRoom)
	router.HandlerFunc(http.MethodGet, "/v1/rooms", app.getAllRooms)
	router.HandlerFunc(http.MethodGet, "/v1/rooms/:booking_id", app.getAllRoomsByBooking)
	
	//router.HandlerFunc(http.MethodPost,"/v1/admin/editroom", app.editRoom)
	router.HandlerFunc(http.MethodGet,"/v1/admin/deletebookingroom/:id", app.deleteRoom)
//-----------------------------------------------------------------
	router.HandlerFunc(http.MethodGet, "/v1/bookings", app.getAllBookings)
//-----------------------------------------------------------------
	router.HandlerFunc(http.MethodGet, "/v1/address/:id", app.getOneRoomAddress)
	router.HandlerFunc(http.MethodGet, "/v1/addresses", app.getAllRoomAddresses)
	router.HandlerFunc(http.MethodGet, "/v1/addresses/:room_id", app.getAllAddressesByRoom)
	
	router.HandlerFunc(http.MethodGet,"/v1/admin/deleteroomaddress/:id", app.deleteAddress)
//-----------------------------------------------------------------
	router.HandlerFunc(http.MethodGet, "/v1/equipment/:id", app.getOneRoomEquipment)
	router.HandlerFunc(http.MethodGet, "/v1/equipments", app.getAllRoomEquipments)
	router.HandlerFunc(http.MethodGet, "/v1/equipments/:room_id", app.getAllEquipmentsByRoom)
	
	router.HandlerFunc(http.MethodGet,"/v1/admin/deleteroomequipment/:id", app.deleteEquipment)
//-----------------------------------------------------------------
	router.HandlerFunc(http.MethodGet, "/v1/organizer/:id", app.getOneBookingOrganizer)
	router.HandlerFunc(http.MethodGet, "/v1/organizers", app.getAllBookingOrganizers)
	router.HandlerFunc(http.MethodGet, "/v1/organizers/:booking_id", app.getAllOrganizersByBooking)
	
	router.HandlerFunc(http.MethodGet,"/v1/admin/deletebookingorganizer/:id", app.deleteOrganizer)
//-----------------------------------------------------------------
    router.HandlerFunc(http.MethodGet, "/v1/bookingparticipant/:id", app.getOneBookingParticipant)
	router.HandlerFunc(http.MethodGet, "/v1/bookingparticipants", app.getAllBookingParticipants)
	router.HandlerFunc(http.MethodGet, "/v1/bookingparticipants/:booking_id", app.getAllParticipantsByBooking)
	
	router.HandlerFunc(http.MethodGet,"/v1/admin/deletebookingparticipant/:id", app.deleteParticipantBookings)
//-----------------------------------------------------------------
    router.HandlerFunc(http.MethodGet, "/v1/organizerparticipant/:id", app.getOneOrganizerParticipant)
	router.HandlerFunc(http.MethodGet, "/v1/organizerparticipants", app.getAllOrganizerParticipants)
	router.HandlerFunc(http.MethodGet, "/v1/organizerparticipants/:organizer_id", app.getAllParticipantsByOrganizer)
	
	router.HandlerFunc(http.MethodGet,"/v1/admin/deleteorganizerparticipant/:id", app.deleteParticipantOrganizers)

	// router.POST("/v1/admin/editroom", app.wrap(secure.ThenFunc(app.editRoom)))
	// router.GET("/v1/admin/deleteroom/:id", app.wrap(secure.ThenFunc(app.deleteRoom)))

	return app.enableCORS(router)
}