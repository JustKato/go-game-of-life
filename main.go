package main

import (
	"fmt"
	"gol2/lib/game"
	"time"
)

func main() {
	g := game.NewGame(32, 32, true)

	for i := 0; i < 300; i++ {
		// Step to the next frame
		g.Step()
		// Clear scree and print the world
		fmt.Print("\x0c", g) // Clear screen and print field.
		// Wait a little
		time.Sleep(time.Millisecond)
	}
}
