package main

import (
	"log"
	"net/http"

	poker "github.com/Shang47/simple-web-server"
)

func main() {
	store, close, err := poker.NewMariaPlayerStore()

	if err != nil {
		log.Fatalf("problem creating DB handle %v", err)
	}
	defer close()

	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.Alerter), store)

	server, err := poker.NewPlayerServer(store, game)

	if err != nil {
		log.Fatalf("problem creating player server %v", err)
	}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}

}
