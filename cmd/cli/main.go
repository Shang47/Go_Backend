package main

import (
	"fmt"
	"os"

	poker "github.com/Shang47/simple-web-server"
)

func main() {
	fmt.Println("Let's play poker.")
	fmt.Println("Type {Name} wins to record a win")

	store := poker.NewMariaPlayerStore()

	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
