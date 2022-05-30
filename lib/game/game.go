package game

import (
	"bytes"
	"gol2/lib/world"
	"math/rand"
)

type Game struct {
	// The GameWorld has two snapshots, as to not disturb one another when calculating the cells.
	aWorld, bWorld *world.World
	// The width and height of the two world copies
	w uint32
	h uint32
	// The amount of generatiosn that have been ran
	generations uint64
}

// Generate a new Game of Life session
func NewGame(w uint32, h uint32, doRandomize bool) *Game {

	// Initialize the world
	aWorld := world.NewWorld(w, h)
	if doRandomize {
		// Give life to a couple of random cells.
		for i := 0; uint32(i) < (w * h / 5); i++ {
			// Set a random spot on the world as alive
			aWorld.Set(uint32(rand.Intn(int(w))), uint32(rand.Intn(int(h))), 3)
		}
	}

	// Generate the game world
	return &Game{
		// Generate the game world
		aWorld: aWorld,
		bWorld: world.NewWorld(w, h),
		// Start the generation step at 0
		generations: 0,
		w:           w,
		h:           h,
	}

}

// Proceed to the next step of the game world
func (g *Game) Step() {

	var y uint32
	var x uint32

	// Go through the whole world
	for y = 0; y < g.h; y++ {
		for x = 0; x < g.w; x++ {
			g.bWorld.Set(x, y, g.aWorld.ProcessCell(int(x), int(y)))
		}
	}

	// Swap the two worlds
	g.aWorld, g.bWorld = g.bWorld, g.aWorld
	// Add to the amount of generations that have ran
	g.generations++
}

// String returns the game board as a string.
func (g *Game) String() string {
	var buf bytes.Buffer
	var y uint32
	var x uint32

	for y = 0; y < g.h; y++ {
		for x = 0; x < g.w; x++ {
			b := byte(' ')

			if g.aWorld.IsAlive(int32(x), int32(y)) {
				b = '*'
			}

			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}

	return buf.String()
}
