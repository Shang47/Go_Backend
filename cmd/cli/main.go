package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/Shang47/simple-web-server"
)

func main() {
	store, close, err := poker.NewMariaPlayerStore()

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play poker.")
	fmt.Println("Type {Name} wins to record a win")

	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
