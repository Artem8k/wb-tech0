package database

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	Client *sqlx.DB
}

func MustRun() *Database {
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("PORT")
		user     = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASS")
		dbname   = os.Getenv("DB_NAME")
	)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s",
		host, port, user, password, dbname)

	db, err := sqlx.Open("pgx", psqlInfo)

	if err != nil {
		panic(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		panic(pingErr)
	}

	fmt.Println("Successfully pinnged!")
	fmt.Println("Successfully connected!")

	return &Database{
		Client: db,
	}
}
