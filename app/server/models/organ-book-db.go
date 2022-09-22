package models

import (
	"context"
	//"database/sql"
	"fmt"
	"time"
)


func (m *DBModel) GetBookingOrganizer(id int) (*Organizer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, organizer_name, organizer_email, created_at, updated_at
						FROM organizers 
						WHERE id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var organizer Organizer

	err := row.Scan(
		&organizer.ID,
		&organizer.OrganizerName,
		&organizer.OrganizerEmail,
		&organizer.CreatedAt,
		&organizer.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	query = `SELECT mg.id, mg.organizer_id, mg.booking_id, g.booking_name
					 FROM organizers_bookings mg
					 LEFT JOIN bookings g on (g.id = mg.booking_id)
					 WHERE mg.organizer_id = $1`

	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	bookings := make(map[int]string)
	for rows.Next() {
		var mg OrganizerBooking
		err := rows.Scan(
			&mg.ID,
			&mg.OrganizerID,
			&mg.BookingID,
			&mg.Booking.BookingName,
		) 

		if err != nil {
			return nil, err
		}

		bookings[mg.ID] = mg.Booking.BookingName
	}

	organizer.OrganizerBooking = bookings

	return &organizer, nil
}

func (m *DBModel) AllBookingOrganizer(booking ...int) ([]*Organizer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	where := ""

	if len (booking) > 0 {
		where = fmt.Sprintf("WHERE id in (SELECT organizer_id FROM organizers_bookings WHERE booking_id = %d)", booking[0])
	}

	query := fmt.Sprintf(`SELECT id, organizer_name, organizer_email, created_at, updated_at 
	FROM organizers %s
	ORDER BY organizer_name`, where)

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var organizers []*Organizer

	for rows.Next(){
			var organizer Organizer
			err := rows.Scan(
				&organizer.ID,
				&organizer.OrganizerName,
				&organizer.OrganizerEmail,
				&organizer.CreatedAt,
				&organizer.UpdatedAt,
				
			)

			if err != nil {
				return nil, err
			}

			bookingQuery := `SELECT mg.id, mg.organizer_id, mg.booking_id, g.booking_name
					 FROM organizers_bookings mg
					 LEFT JOIN bookings g on (g.id = mg.booking_id)
					 WHERE mg.organizer_id = $1`

			bookingRows, _ := m.DB.QueryContext(ctx, bookingQuery, organizer.ID)


			bookings := make(map[int]string)
			for bookingRows.Next() {
				var mg OrganizerBooking
				err := bookingRows.Scan(
					&mg.ID,
					&mg.OrganizerID,
					&mg.BookingID,
					&mg.Booking.BookingName,
				) 

				if err != nil {
					return nil, err
				}

				bookings[mg.ID] = mg.Booking.BookingName
			}

			bookingRows.Close()

			organizer.OrganizerBooking = bookings
			organizers = append(organizers, &organizer)

	}

	return organizers, nil
}

// func (m *DBModel) OrganizerBookingsAll() ([]*Booking, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	query := `SELECT id, booking_name, desc_booking, date_booking, from_booking, till_booking, regular, notification, canceled, created_at, updated_at
// 	FROM bookings
// 	ORDER BY booking_name`

// 	rows, err := m.DB.QueryContext(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var bookings []*Booking
	
// 	for rows.Next() {
// 		var g Booking
// 		err := rows.Scan(
// 			&g.ID,
// 			&g.BookingName,
// 			&g.DescBooking,
// 			&g.DateBooking,
// 			&g.FromBooking,
// 			&g.TillBooking,
// 			&g.Regular,
// 			&g.Notification,
// 			&g.Canceled,
// 			&g.CreatedAt,
// 			&g.UpdatedAt,
// 		)

// 		if err != nil {
// 			return nil, err
// 		}

// 		bookings = append(bookings, &g)
// 	}

// 	return bookings, nil
// }

func (m *DBModel) InsertOrganizer(organizer Organizer) error  {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt :=  `INSERT INTO organizers (organizer_name, organizer_email, created_at, updated_at)
						VALUES ($1, $2, $3, $4)`

	_, err := m.DB.ExecContext(ctx, stmt, 
		organizer.OrganizerName,
		organizer.OrganizerEmail,
		organizer.CreatedAt,
		organizer.UpdatedAt,
		
		)

		if err != nil {
			return err
		}

		return nil
}



func (m *DBModel) UpdateOrganizer(organizer Organizer) error  {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt :=  `UPDATE ORGANIZERS SET organizer_name = $1, organizer_email = $2, updated_at = $3
						WHERE id = $4`

	_, err := m.DB.ExecContext(ctx, stmt, 
		organizer.OrganizerName,
		organizer.OrganizerEmail,
		organizer.UpdatedAt,
		organizer.ID,
		)

		if err != nil {
			return err
		}

		return nil
}


func (m *DBModel) DeleteOrganizer (id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `delete from organizers_bookings where organizer_id = $1`

	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil

}
