package main

import (
	"log"
	"net/http"
	"time"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Function called")
	time.After(6 * time.Second)
	// select {
	// case <-time.After(6 * time.Second):
	// 	fmt.Println("received timeout")
	// case <-r.Context().Done():
	// 	fmt.Println("received cancellation")
	// 	// return
	// }

	w.Write([]byte("HELLO"))
}
