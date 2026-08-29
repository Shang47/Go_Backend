package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewMariaPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", server))
	defer server.store.close()
}
