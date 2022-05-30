package main

import (
	"fmt"
	"gol2/lib/game"
	"image/color"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// The main game of the screen
var g *game.Game
var runSpeed float32 = 2
var runGame bool = false

func runGameLogic() {

	// Start ticking the game
	for true {
		// Only run the game if the runGame variable is true
		if runGame {
			// Step to the next frame
			g.Step()
			time.Sleep(time.Second * time.Duration(runSpeed))
		}
	}

}

var bgColor = color.RGBA{R: 26, G: 26, B: 26, A: 255}
var fgColor = color.RGBA{R: 228, G: 219, B: 210, A: 255}
var cellSize int = 16

var xOffset, yOffset int = 0, 0

func drawGame() {

	// get the middle of the screen
	width, height := g.GetSize()

	// Get the current state of the world
	worldState := g.GetCurrentWorldState()

	// Figure out the cell size in pixels

	// Find the center of the screen
	hCenter := rl.GetScreenWidth() / 2
	vCenter := rl.GetScreenHeight() / 2

	// Subtract the half of our grid's size from the center
	hCenter -= int((width * uint32(cellSize)) / 2)
	vCenter -= int((height * uint32(cellSize)) / 2)

	// Handle offsets
	hCenter -= (xOffset * int(width/3))
	vCenter -= (yOffset * int(height/3))

	for y := 0; y < int(height); y++ {
		for x := 0; x < int(width); x++ {

			xPos := hCenter + (x * cellSize)
			yPos := vCenter + (y * cellSize)

			if worldState[x][y] != 0 {
				// Alive
				rl.DrawRectangle(int32(xPos), int32(yPos), int32(cellSize), int32(cellSize), fgColor)
			} else {
				// Dead
				rl.DrawRectangleLines(int32(xPos), int32(yPos), int32(cellSize), int32(cellSize), fgColor)
			}

		}
	}

}

func drawHud() {

	var stateName string = "Not Running"
	if runGame {
		stateName = "Running"
	}

	rl.DrawText(fmt.Sprintf("Game Status: %s", stateName), 10, 10, 20, fgColor)
	rl.DrawText(fmt.Sprintf("Generation: %v", g.GetGeneration()), 10, 30, 20, fgColor)
	rl.DrawText(fmt.Sprintf("Frame Time: %.2f", rl.GetFrameTime()), 10, 50, 20, fgColor)
	rl.DrawText(fmt.Sprintf("Cell Size: %v", cellSize), 10, 70, 20, fgColor)
	rl.DrawText(fmt.Sprintf("Game Speed: %v ( * Second )", runSpeed), 10, 90, 20, fgColor)

	rl.DrawText("To pan around use W, A, S, D or arrow keys", 10, int32(rl.GetScreenHeight())-70, 20, fgColor)
	rl.DrawText("Hold down LEFT-CTRL and scroll to change the game speed", 10, int32(rl.GetScreenHeight())-50, 20, fgColor)
	rl.DrawText("Use your scroll wheel to zoom in/out, press space to pause or play.", 10, int32(rl.GetScreenHeight())-30, 20, fgColor)

}

func handleControls() {

	if !rl.IsKeyDown(rl.KeyLeftControl) {
		cellSize += int(rl.GetMouseWheelMove())

		if cellSize < 6 {
			cellSize = 6
		}

		if cellSize > 128 {
			cellSize = 128
		}
	} else if rl.IsKeyDown(rl.KeyLeftControl) {
		runSpeed += .25 * runSpeed * float32(rl.GetMouseWheelMove())
		if runSpeed <= 0.05 {
			runSpeed = 0.05
		}
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		runGame = !runGame
	}

	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		// Up
		yOffset -= 1
	}

	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		// Down
		yOffset += 1
	}

	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		// Left
		xOffset -= 1
	}

	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		// Right
		xOffset += 1
	}

}

func main() {
	// Initialize the game
	g = game.NewGame(128, 128, true)

	// Run the game logic ticker on a separate thread
	go runGameLogic()

	// Initialize the raylib window
	rl.InitWindow(1280, 720, "Golang Game of Life")

	rl.SetWindowState(rl.FlagWindowResizable)

	// Set the target FPS of the window
	rl.SetTargetFPS(int32(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor())))

	// Run the draw logic
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(bgColor)

		// Handle Controls
		handleControls()
		// Draw the game
		drawGame()
		// Draw the game hud
		drawHud()

		rl.EndDrawing()
	}

}
