package models

import (
	"context"
	//"database/sql"
	"fmt"
	"time"
)

func (m *DBModel) GetRoomAddress(id int) (*Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, city, street, building, floor, created_at, updated_at
						FROM addresses 
						WHERE id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var address Address

	err := row.Scan(
		&address.ID,
		&address.City,
		&address.Street,
		&address.Building,
		&address.Floor,
		&address.CreatedAt,
		&address.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	query = `SELECT mg.id, mg.address_id, mg.room_id, g.room_name
					 FROM addresses_rooms mg
					 LEFT JOIN rooms g on (g.id = mg.room_id)
					 WHERE mg.address_id = $1`

	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	rooms := make(map[int]string)
	for rows.Next() {
		var mg AddressRoom
		err := rows.Scan(
			&mg.ID,
			&mg.AddressID,
			&mg.RoomID,
			&mg.Room.RoomName,
		) 

		if err != nil {
			return nil, err
		}

		rooms[mg.ID] = mg.Room.RoomName
	}

	address.AddressRoom = rooms

	return &address, nil
}

func (m *DBModel) AllRoomAddress(room ...int) ([]*Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	where := ""

	if len (room) > 0 {
		where = fmt.Sprintf("WHERE id in (SELECT address_id FROM addresses_rooms WHERE room_id = %d)", room[0])
	}

	query := fmt.Sprintf(`SELECT id, city, street, building, floor, created_at, updated_at 
	FROM addresses %s
	ORDER BY city`, where)

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var addresses []*Address

	for rows.Next(){
			var address Address
			err := rows.Scan(
				&address.ID,
				&address.City,
				&address.Street,
				&address.Building,
				&address.Floor,
				&address.CreatedAt,
				&address.UpdatedAt,
				
			)

			if err != nil {
				return nil, err
			}

			roomQuery := `SELECT mg.id, mg.address_id, mg.room_id, g.room_name
					 FROM addresses_rooms mg
					 LEFT JOIN rooms g on (g.id = mg.room_id)
					 WHERE mg.address_id = $1`

			roomRows, _ := m.DB.QueryContext(ctx, roomQuery, address.ID)


			rooms := make(map[int]string)
			for roomRows.Next() {
				var mg AddressRoom
				err := roomRows.Scan(
					&mg.ID,
					&mg.AddressID,
					&mg.RoomID,
					&mg.Room.RoomName,
				) 

				if err != nil {
					return nil, err
				}

				rooms[mg.ID] = mg.Room.RoomName
			}

			roomRows.Close()

			address.AddressRoom = rooms
			addresses = append(addresses, &address)

	}

	return addresses, nil
}

func (m *DBModel) RoomsAll() ([]*Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, room_name, space, security_email, visible, created_at, updated_at
	FROM rooms
	ORDER BY room_name`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []*Room
	
	for rows.Next() {
		var g Room
		err := rows.Scan(
			&g.ID,
			&g.RoomName,
			&g.Space,
			&g.SecurityEmail,
			&g.Visible,
			&g.CreatedAt,
			&g.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		rooms = append(rooms, &g)
	}

	return rooms, nil
}

func (m *DBModel) InsertAddress(address Address) error  {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt :=  `INSERT INTO addresses (city, street, building, floor, created_at, updated_at)
						VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := m.DB.ExecContext(ctx, stmt, 
		address.City,
		address.Street,
		address.Building,
		address.Floor,
		address.CreatedAt,
		address.UpdatedAt,
		
		)

		if err != nil {
			return err
		}

		return nil
}



func (m *DBModel) UpdateAddress(address Address) error  {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt :=  `UPDATE ADDRESSES SET city = $1, street = $2, building = $3, floor = $4, updated_at = $5
						WHERE id = $6`

	_, err := m.DB.ExecContext(ctx, stmt, 
		address.City,
		address.Street,
		address.Building,
		address.Floor,
		address.UpdatedAt,
		address.ID,
		)

		if err != nil {
			return err
		}

		return nil
}


func (m *DBModel) DeleteAddress (id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `delete from addresses_rooms where address_id = $1`

	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil

}
