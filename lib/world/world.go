package world

// A world is a current snapshot of a Game, it contains all of the cells.
type World struct {
	cells [][]uint8

	// The width of the world
	w uint32
	// The height of the world
	h uint32
}

// NewWorld creates a new Game of Life playing field, this is not the game itself, it's just something that keeps track of the world state.
func NewWorld(w uint32, h uint32) *World {

	// Spawn the grid's first column
	cells := make([][]uint8, h)
	// Go through all of the rows of width 0
	for i := range cells {
		// Set the row size to be the actual size
		cells[i] = make([]uint8, w)
	}

	// Return the newly generated world
	return &World{
		cells: cells,
		w:     w,
		h:     h,
	}

}

// Set the value of a cell
func (w World) Set(x uint32, y uint32, v uint8) {
	// Set the value
	w.cells[x][y] = v
}

func (w World) IsAlive(x int32, y int32) bool {
	return w.Get(x, y) != 0
}

// Get the current state of the cell
func (w World) Get(x int32, y int32) uint8 {
	// Add the width of the world
	x += int32(w.w)
	// Modulus to wrap around
	x %= int32(w.w)

	// Add in the height of the world
	y += int32(w.h)
	// Modulus to wrap around the world
	y %= int32(w.h)
	// Return the value
	return w.cells[y][x]
}

// ProcessCell returns the future state of the specified cell
func (w *World) ProcessCell(x int, y int) uint8 {
	// The amount of cells that are alive around us
	alive := 0

	// Go through all our neighbors
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && w.IsAlive(int32(x+i), int32(y+j)) {
				alive++
			}
		}
	}

	// 3 Neighbors = revive/stay on
	// 2 Neighbors = Maintain current state ( dead/alive )
	// Anything else = Death

	if alive == 3 || (alive == 2 && w.IsAlive(int32(x), int32(y))) {
		// The cell is alive, and we will reflect the number of neighbors onto it
		return uint8(alive)
	} else {
		// The cell is dead
		return 0
	}
}

func (w *World) GetCurrentState() [][]uint8 {
	return w.cells
}
