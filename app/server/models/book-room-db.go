package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) GetBookingRoom(id int) (*Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, room_name, space, security_email, visible, created_at, updated_at
						FROM rooms 
						WHERE id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var room Room

	err := row.Scan(
		&room.ID,
		&room.RoomName,
		&room.Space,
		&room.SecurityEmail,
		&room.Visible,
		&room.CreatedAt,
		&room.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	query = `SELECT mg.id, mg.room_id, mg.booking_id, g.booking_name
					 FROM rooms_bookings mg
					 LEFT JOIN bookings g on (g.id = mg.booking_id)
					 WHERE mg.room_id = $1`

	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	bookings := make(map[int]string)
	for rows.Next() {
		var mg RoomBooking
		err := rows.Scan(
			&mg.ID,
			&mg.RoomID,
			&mg.BookingID,
			&mg.Booking.BookingName,
		) 

		if err != nil {
			return nil, err
		}

		bookings[mg.ID] = mg.Booking.BookingName
	}

	room.RoomBooking = bookings

	return &room, nil
}

func (m *DBModel) AllBookingRoom(booking ...int) ([]*Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	where := ""

	if len (booking) > 0 {
		where = fmt.Sprintf("WHERE id in (SELECT room_id FROM rooms_bookings WHERE booking_id = %d)", booking[0])
	}

	query := fmt.Sprintf(`SELECT id, room_name, space, security_email, visible, created_at, updated_at 
	FROM rooms %s
	ORDER BY room_name`, where)

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rooms []*Room

	for rows.Next(){
			var room Room
			err := rows.Scan(
				&room.ID,
				&room.RoomName,
				&room.Space,
				&room.SecurityEmail,
				&room.Visible,
				&room.CreatedAt,
				&room.UpdatedAt,
				
			)

			if err != nil {
				return nil, err
			}

			bookingQuery := `SELECT mg.id, mg.room_id, mg.booking_id, g.booking_name
					 FROM rooms_bookings mg
					 LEFT JOIN bookings g on (g.id = mg.booking_id)
					 WHERE mg.room_id = $1`

			bookingRows, _ := m.DB.QueryContext(ctx, bookingQuery, room.ID)


			bookings := make(map[int]string)
			for bookingRows.Next() {
				var mg RoomBooking
				err := bookingRows.Scan(
					&mg.ID,
					&mg.RoomID,
					&mg.BookingID,
					&mg.Booking.BookingName,
				) 

				if err != nil {
					return nil, err
				}

				bookings[mg.ID] = mg.Booking.BookingName
			}

			bookingRows.Close()

			room.RoomBooking = bookings
			rooms = append(rooms, &room)

	}

	return rooms, nil
}

func (m *DBModel) BookingsAll() ([]*Booking, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, booking_name, desc_booking, date_booking, from_booking, till_booking, regular, notification, canceled, created_at, updated_at
	FROM bookings
	ORDER BY booking_name`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []*Booking
	
	for rows.Next() {
		var g Booking
		err := rows.Scan(
			&g.ID,
			&g.BookingName,
			&g.DescBooking,
			&g.DateBooking,
			&g.FromBooking,
			&g.TillBooking,
			&g.Regular,
			&g.Notification,
			&g.Canceled,
			&g.CreatedAt,
			&g.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		bookings = append(bookings, &g)
	}

	return bookings, nil
}

func (m *DBModel) InsertRoom(room Room) error  {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt :=  `INSERT INTO rooms (room_name, space, security_email, visible, created_at, updated_at)
						VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := m.DB.ExecContext(ctx, stmt, 
		room.RoomName,
		room.Space,
		room.SecurityEmail,
		room.Visible,
		room.CreatedAt,
		room.UpdatedAt,
		
		)

		if err != nil {
			return err
		}

		return nil
}



func (m *DBModel) UpdateRoom(room Room) error  {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt :=  `UPDATE ROOMS SET room_name = $1, space = $2, security_email = $3, visible = $4, updated_at = $5
						WHERE id = $6`

	_, err := m.DB.ExecContext(ctx, stmt, 
		room.RoomName,
		room.Space,
		room.SecurityEmail,
		room.Visible,
		room.UpdatedAt,
		room.ID,
		)

		if err != nil {
			return err
		}

		return nil
}


func (m *DBModel) DeleteRoom (id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `delete from rooms_bookings where room_id = $1`

	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil

}
