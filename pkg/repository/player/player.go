package player

import (
	"BogeyTrain/pkg/models"
	"database/sql"
)

func GetAllPlayers(db *sql.DB) ([]*models.Player, error) {
	query := `select id, first_name, last_name from players`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
		}
	}(rows)

	var persons []*models.Player
	for rows.Next() {
		var p models.Player
		err := rows.Scan(&p.ID, &p.FirstName, &p.LastName)
		if err != nil {
			return nil, err
		}

		persons = append(persons, &p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return persons, nil
}