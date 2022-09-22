package models

import (
	"context"
	//"database/sql"
	"fmt"
	"time"
)


func (m *DBModel) GetOrganizerParticipant(id int) (*Participant, error) {
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

	query = `SELECT mg.id, mg.participant_id, mg.organizer_id, g.organizer_name
					 FROM participants_organizers mg
					 LEFT JOIN organizers g on (g.id = mg.organizer_id)
					 WHERE mg.participant_id = $1`

	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	organizers := make(map[int]string)
	for rows.Next() {
		var mg ParticipantOrganizer
		err := rows.Scan(
			&mg.ID,
			&mg.ParticipantID,
			&mg.OrganizerID,
			&mg.Organizer.OrganizerName,
		) 

		if err != nil {
			return nil, err
		}

		organizers[mg.ID] = mg.Organizer.OrganizerName
	}

	participant.ParticipantOrganizer = organizers

	return &participant, nil
}

func (m *DBModel) AllOrganizerParticipant(organizer ...int) ([]*Participant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	where := ""

	if len (organizer) > 0 {
		where = fmt.Sprintf("WHERE id in (SELECT participant_id FROM participants_organizers WHERE organizer_id = %d)", organizer[0])
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

			organizerQuery := `SELECT mg.id, mg.participant_id, mg.organizer_id, g.organizer_name
					 FROM participants_organizers mg
					 LEFT JOIN organizers g on (g.id = mg.organizer_id)
					 WHERE mg.participant_id = $1`

			organizerRows, _ := m.DB.QueryContext(ctx, organizerQuery, participant.ID)


			organizers := make(map[int]string)
			for organizerRows.Next() {
				var mg ParticipantOrganizer
				err := organizerRows.Scan(
					&mg.ID,
					&mg.ParticipantID,
					&mg.OrganizerID,
					&mg.Organizer.OrganizerName,
				) 

				if err != nil {
					return nil, err
				}

				organizers[mg.ID] = mg.Organizer.OrganizerName
			}

			organizerRows.Close()

			participant.ParticipantOrganizer = organizers
			participants = append(participants, &participant)

	}

	return participants, nil
}

func (m *DBModel) OrganizersAll() ([]*Organizer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, organizer_name, organizer_email, created_at, updated_at
	FROM organizers
	ORDER BY organizer_name`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var organizers []*Organizer
	
	for rows.Next() {
		var g Organizer
		err := rows.Scan(
			&g.ID,
			&g.OrganizerName,
			&g.OrganizerEmail,
			&g.CreatedAt,
			&g.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		organizers = append(organizers, &g)
	}

	return organizers, nil
}

func (m *DBModel) InsertParticipant(participant Participant) error  {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt :=  `INSERT INTO participants (participant_name, participant_email, participant_car, created_at, updated_at)
						VALUES ($1, $2, $3, $4, $5)`

	_, err := m.DB.ExecContext(ctx, stmt, 
		participant.ParticipantName,
		participant.ParticipantEmail,
		participant.ParticipantCar,
		participant.CreatedAt,
		participant.UpdatedAt,
		
		)

		if err != nil {
			return err
		}

		return nil
}



func (m *DBModel) UpdateParticipant(participant Participant) error  {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt :=  `UPDATE PARTICIPANTS SET participant_name = $1, participant_email = $2, participant_car = $3, updated_at = $4
						WHERE id = $5`

	_, err := m.DB.ExecContext(ctx, stmt, 
		participant.ParticipantName,
		participant.ParticipantEmail,
		participant.ParticipantCar,
		participant.UpdatedAt,
		participant.ID,
		)

		if err != nil {
			return err
		}

		return nil
}


func (m *DBModel) DeleteParticipantOrganizers (id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `delete from participants_organizers where participant_id = $1`

	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil

}
