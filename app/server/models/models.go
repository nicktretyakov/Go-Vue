package models

import (
	"database/sql"
	"time"
)

type Models struct {
	DB DBModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type Room struct {
	ID            int            `json:"id"`
	RoomName      string         `json:"room_name"`
	Space         int            `json:"space"`
	SecurityEmail string         `json:"security_email"`
	Visible       bool           `json:"visible"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	RoomBooking   map[int]string `json:"bookings"`
}

type Booking struct {
	ID           int       `json:"id"`
	BookingName  string    `json:"booking_name"`
	DescBooking  string    `json:"desc_booking"`
	DateBooking  time.Time `json:"date_booking"`
	FromBooking  time.Time `json:"from_booking"`
	TillBooking  time.Time `json:"till_booking"`
	Regular      bool      `json:"regular"`
	Notification int       `json:"notification"`
	Canceled     bool      `json:"canceled"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type RoomBooking struct {
	ID        int       `json:"-"`
	RoomID    int       `json:"-"`
	BookingID int       `json:"-"`
	Booking   Booking   `json:"booking"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// --------------------------------
type Participant struct {
	ID                   int            `json:"id"`
	ParticipantName      string         `json:"participant_name"`
	ParticipantEmail     string         `json:"participant_email"`
	ParticipantCar       string         `json:"participant_car"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
	ParticipantOrganizer map[int]string `json:"organizer"`
	ParticipantBooking   map[int]string `json:"booking"`
}

type Organizer struct {
	ID             int       `json:"id"`
	OrganizerName  string    `json:"organizer_name"`
	OrganizerEmail string    `json:"organizer_email"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	OrganizerBooking map[int]string `json:"booking"`
}

type ParticipantOrganizer struct {
	ID            int       `json:"-"`
	ParticipantID int       `json:"-"`
	OrganizerID   int       `json:"-"`
	Organizer     Organizer `json:"organizer"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}

// -----------------------------------------
type OrganizerBooking struct {
	ID          int       `json:"-"`
	OrganizerID int       `json:"-"`
	BookingID   int       `json:"-"`
	Booking     Booking   `json:"booking"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

// -----------------------------------------
type ParticipantBooking struct {
	ID            int       `json:"-"`
	ParticipantID int       `json:"-"`
	BookingID     int       `json:"-"`
	Booking       Booking   `json:"booking"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}
//------------------------------------------
type Address struct {
	ID            int            `json:"id"`
	City          string         `json:"city"`
	Street        string         `json:"street"`
	Building      string         `json:"building"`
	Floor         string         `json:"floor"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	AddressRoom   map[int]string `json:"rooms"`
}

type AddressRoom struct {
	ID         int       `json:"-"`
	AddressID  int       `json:"-"`
	RoomID     int       `json:"-"`
	Room       Room      `json:"room"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
}
//---------------------------------------
type Equipment struct {
	ID            int            `json:"id"`
	EquipmentName string         `json:"equipment_name"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	EquipmentRoom   map[int]string `json:"rooms"`
}

type EquipmentRoom struct {
	ID         int       `json:"-"`
	EquipmentID  int     `json:"-"`
	RoomID     int       `json:"-"`
	Room       Room      `json:"room"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
}

//-----------------------------------
type User struct {
	ID       int
	Email    string
	Password string
}
