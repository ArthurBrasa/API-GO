package db

import (
	"database/sql"
	"fmt"

	"github.com/ArthurBrasa/API-GO/configs"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	// Pegar as configs de conexao com o DB
	config := configs.GetDB()

	string_connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Pass, config.Database)

	connection, err := sql.Open("postgres", string_connection)
	if err != nil {
		panic(err)
	}

	err = connection.Ping()

	return connection, err
}
