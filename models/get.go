package models

import "github.com/ArthurBrasa/API-GO/db"

func Get(id int64) (todo Todo, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer connection.Close()

	row := connection.QueryRow(`SELECT * FROM todos WHERE id = $1`, id)

	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}
