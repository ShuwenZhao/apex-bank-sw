package main

import (
	"database/sql"
	"log"

	"github.com/ShuwenZhao/apex-bank-sw/api"
	db "github.com/ShuwenZhao/apex-bank-sw/db/sqlc"
	"github.com/ShuwenZhao/apex-bank-sw/util"
	_ "github.com/golang/mock/mockgen/model"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
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

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
