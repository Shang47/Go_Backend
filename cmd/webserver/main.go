package main

import (
	"log"
	"net/http"

	poker "github.com/Shang47/simple-web-server"
)

func main() {
	server := poker.NewPlayerServer(poker.NewMariaPlayerStore())

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}

	//defer server.store.close()
}
