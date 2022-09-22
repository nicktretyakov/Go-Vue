package models

import (
	"context"
	//"database/sql"
	"fmt"
	"time"
)

func (m *DBModel) GetRoomEquipment(id int) (*Equipment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, equipment_name, created_at, updated_at
						FROM equipments 
						WHERE id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var equipment Equipment

	err := row.Scan(
		&equipment.ID,
		&equipment.EquipmentName,
		&equipment.CreatedAt,
		&equipment.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	query = `SELECT mg.id, mg.equipment_id, mg.room_id, g.room_name
					 FROM equipments_rooms mg
					 LEFT JOIN rooms g on (g.id = mg.room_id)
					 WHERE mg.equipment_id = $1`

	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	rooms := make(map[int]string)
	for rows.Next() {
		var mg EquipmentRoom
		err := rows.Scan(
			&mg.ID,
			&mg.EquipmentID,
			&mg.RoomID,
			&mg.Room.RoomName,
		) 

		if err != nil {
			return nil, err
		}

		rooms[mg.ID] = mg.Room.RoomName
	}

	equipment.EquipmentRoom = rooms

	return &equipment, nil
}

func (m *DBModel) AllRoomEquipment(room ...int) ([]*Equipment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	where := ""

	if len (room) > 0 {
		where = fmt.Sprintf("WHERE id in (SELECT equipment_id FROM equipments_rooms WHERE room_id = %d)", room[0])
	}

	query := fmt.Sprintf(`SELECT id, equipment_name, created_at, updated_at 
	FROM equipments %s
	ORDER BY equipment_name`, where)

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var equipments []*Equipment

	for rows.Next(){
			var equipment Equipment
			err := rows.Scan(
				&equipment.ID,
				&equipment.EquipmentName,
				&equipment.CreatedAt,
				&equipment.UpdatedAt,
				
			)

			if err != nil {
				return nil, err
			}

			roomQuery := `SELECT mg.id, mg.equipment_id, mg.room_id, g.room_name
					 FROM equipments_rooms mg
					 LEFT JOIN rooms g on (g.id = mg.room_id)
					 WHERE mg.equipment_id = $1`

			roomRows, _ := m.DB.QueryContext(ctx, roomQuery, equipment.ID)


			rooms := make(map[int]string)
			for roomRows.Next() {
				var mg EquipmentRoom
				err := roomRows.Scan(
					&mg.ID,
					&mg.EquipmentID,
					&mg.RoomID,
					&mg.Room.RoomName,
				) 

				if err != nil {
					return nil, err
				}

				rooms[mg.ID] = mg.Room.RoomName
			}

			roomRows.Close()

			equipment.EquipmentRoom = rooms
			equipments = append(equipments, &equipment)

	}

	return equipments, nil
}

// func (m *DBModel) RoomsAll() ([]*Room, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	query := `SELECT id, room_name, space, security_email, visible, created_at, updated_at
// 	FROM rooms
// 	ORDER BY room_name`

// 	rows, err := m.DB.QueryContext(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var rooms []*Room
	
// 	for rows.Next() {
// 		var g Room
// 		err := rows.Scan(
// 			&g.ID,
// 			&g.RoomName,
// 			&g.Space,
// 			&g.SecurityEmail,
// 			&g.Visible,
// 			&g.CreatedAt,
// 			&g.UpdatedAt,
// 		)

// 		if err != nil {
// 			return nil, err
// 		}

// 		rooms = append(rooms, &g)
// 	}

// 	return rooms, nil
// }

func (m *DBModel) InsertEquipment(equipment Equipment) error  {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt :=  `INSERT INTO equipments (equipment_name,created_at, updated_at)
						VALUES ($1, $2, $3)`

	_, err := m.DB.ExecContext(ctx, stmt, 
		equipment.EquipmentName,
		equipment.CreatedAt,
		equipment.UpdatedAt,
		
		)

		if err != nil {
			return err
		}

		return nil
}



func (m *DBModel) UpdateEquipment(equipment Equipment) error  {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt :=  `UPDATE EQUIPMENTS SET equipment_name = $1, updated_at = $2
						WHERE id = $3`

	_, err := m.DB.ExecContext(ctx, stmt, 
		equipment.EquipmentName,
		equipment.UpdatedAt,
		equipment.ID,
		)

		if err != nil {
			return err
		}

		return nil
}


func (m *DBModel) DeleteEquipment (id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `delete from equipments_rooms where equipment_id = $1`

	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil

}
