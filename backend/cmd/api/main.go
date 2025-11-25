package main

import (
	"log"

	"github.com/szoumoc/golang+angular/internal/store"
)

func main() {
	store := store.NewPostgresStorage(nil)
	app := &application{
		config: config{
			addr: ":8080",
		},
		store: store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
