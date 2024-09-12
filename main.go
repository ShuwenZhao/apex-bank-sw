package main

import (
	"database/sql"
	"log"

	"github.com/ShuwenZhao/apex-bank-sw/api"
	db "github.com/ShuwenZhao/apex-bank-sw/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	log.Println("Connecting to database with connection string:", dbSource)

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// Ping the database to test the connection
	err = conn.Ping()
	if err != nil {
		log.Fatal("cannot ping db:", err)
	} else {
		log.Println("Successfully connected to the database")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
