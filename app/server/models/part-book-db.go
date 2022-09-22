package models

import (
	"context"
	//"database/sql"
	"fmt"
	"time"
)

func (m *DBModel) GetBookingParticipant(id int) (*Participant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, participant_name, participant_email, participant_car, created_at, updated_at
						FROM participants 
						WHERE id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var participant Participant

	err := row.Scan(
		&participant.ID,
		&participant.ParticipantName,
		&participant.ParticipantEmail,
		&participant.ParticipantCar,
		&participant.CreatedAt,
		&participant.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	query = `SELECT mg.id, mg.participant_id, mg.booking_id, g.booking_name
					 FROM participants_bookings mg
					 LEFT JOIN bookings g on (g.id = mg.booking_id)
					 WHERE mg.participant_id = $1`

	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	bookings := make(map[int]string)
	for rows.Next() {
		var mg ParticipantBooking
		err := rows.Scan(
			&mg.ID,
			&mg.ParticipantID,
			&mg.BookingID,
			&mg.Booking.BookingName,
		) 

		if err != nil {
			return nil, err
		}

		bookings[mg.ID] = mg.Booking.BookingName
	}

	participant.ParticipantBooking = bookings

	return &participant, nil
}

func (m *DBModel) AllBookingParticipant(booking ...int) ([]*Participant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	where := ""

	if len (booking) > 0 {
		where = fmt.Sprintf("WHERE id in (SELECT participant_id FROM participants_bookings WHERE booking_id = %d)", booking[0])
	}

	query := fmt.Sprintf(`SELECT id, participant_name, participant_email, participant_car, created_at, updated_at 
	FROM participants %s
	ORDER BY participant_name`, where)

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var participants []*Participant

	for rows.Next(){
			var participant Participant
			err := rows.Scan(
				&participant.ID,
				&participant.ParticipantName,
				&participant.ParticipantEmail,
				&participant.ParticipantCar,
				&participant.CreatedAt,
				&participant.UpdatedAt,
				
			)

			if err != nil {
				return nil, err
			}

			bookingQuery := `SELECT mg.id, mg.participant_id, mg.booking_id, g.booking_name
					 FROM participants_bookings mg
					 LEFT JOIN bookings g on (g.id = mg.booking_id)
					 WHERE mg.participant_id = $1`

			bookingRows, _ := m.DB.QueryContext(ctx, bookingQuery, participant.ID)


			bookings := make(map[int]string)
			for bookingRows.Next() {
				var mg ParticipantBooking
				err := bookingRows.Scan(
					&mg.ID,
					&mg.ParticipantID,
					&mg.BookingID,
					&mg.Booking.BookingName,
				) 

				if err != nil {
					return nil, err
				}

				bookings[mg.ID] = mg.Booking.BookingName
			}

			bookingRows.Close()

			participant.ParticipantBooking = bookings
			participants = append(participants, &participant)

	}

	return participants, nil
}

// func (m *DBModel) BookingsAll() ([]*Booking, error) {
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

// func (m *DBModel) InsertParticipant(participant Participant) error  {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	stmt :=  `INSERT INTO participants (participant_name, space, security_email, visible, created_at, updated_at)
// 						VALUES ($1, $2, $3, $4, $5, $6)`

// 	_, err := m.DB.ExecContext(ctx, stmt, 
// 		participant.ParticipantName,
// 		participant.Space,
// 		participant.SecurityEmail,
// 		participant.Visible,
// 		participant.CreatedAt,
// 		participant.UpdatedAt,
		
// 		)

// 		if err != nil {
// 			return err
// 		}

// 		return nil
// }



// func (m *DBModel) UpdateParticipant(participant Participant) error  {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	stmt :=  `UPDATE PARTICIPANTS SET participant_name = $1, space = $2, security_email = $3, visible = $4, updated_at = $5
// 						WHERE id = $6`

// 	_, err := m.DB.ExecContext(ctx, stmt, 
// 		participant.ParticipantName,
// 		participant.Space,
// 		participant.SecurityEmail,
// 		participant.Visible,
// 		participant.UpdatedAt,
// 		participant.ID,
// 		)

// 		if err != nil {
// 			return err
// 		}

// 		return nil
// }


 func (m *DBModel) DeleteParticipantBookings (id int) error {
 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
 	defer cancel()

 	stmt := `delete from participants_bookings where participant_id = $1`

 	_, err := m.DB.ExecContext(ctx, stmt, id)
 	if err != nil {
 		return err
 	}

 	return nil

 }
