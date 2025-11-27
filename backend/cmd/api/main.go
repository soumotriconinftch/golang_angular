package main

import (
	"log"

	"github.com/szoumoc/golang+angular/internal/db"
	"github.com/szoumoc/golang+angular/internal/env"
	"github.com/szoumoc/golang+angular/internal/store"
)

func main() {

	dbconf := dbConfig{
		addr:         env.GetString("DB_ADDR", "LOL"),
		maxOpenConns: 30,
		maxIdleConns: 30,
		maxIdleTime:  "15m",
	}

	db, err := db.New(
		dbconf.addr,
		dbconf.maxOpenConns,
		dbconf.maxIdleConns,
		dbconf.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	log.Println("DB connected!")
	store := store.NewPostgresStorage(db)

	app := &application{
		config: config{
			addr: env.GetString("ADDR", ":9000"),
			db:   dbconf,
		},
		store: store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
