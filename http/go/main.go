package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type incomingEvent struct {
	Data interface{} `json:"data"`
}

func main() {
	http.HandleFunc("/greeting", func(rw http.ResponseWriter, r *http.Request) {
		var event incomingEvent
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&event)
		fmt.Println(event.Data)
		fmt.Fprintf(rw, "Hello, World!")
	})
	log.Fatal(http.ListenAndServe(":8088", nil))
}
