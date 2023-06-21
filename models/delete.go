package models

import "github.com/ArthurBrasa/API-GO/db"

func Delete(id int64) (int64, error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer connection.Close()

	res, err := connection.Exec(`DELETE FROM todos WHERE id=$4`, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
