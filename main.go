//Jeu d'automatisation avec un cap de productivité à passer sur chaque niveau !

package main

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const screenWidth = 960
const screenHeight = 540
const windowTitle = "Raylib Go Test"
const targetFPS = 60

type Screen int

const (
	MainMenu Screen = 0
	Level1   Screen = 1
)

var screen Screen = MainMenu

func main() {
	rl.InitWindow(screenWidth, screenHeight, windowTitle)
	defer rl.CloseWindow()

	rl.SetTargetFPS(targetFPS)

	gridSize := 0
	if screenWidth > screenHeight {
		gridSize = screenHeight * 0.85
	} else {
		gridSize = screenWidth * 0.85
	}

	// If int not divisable by 2 then it's even so we add one to fix it
	// Prevent placing grid on a subpixel
	if gridSize%2 != 0 {
		gridSize++
	}

	gridCells := 25
	cellSize := gridSize / gridCells

	if screenWidth < gridSize || screenHeight < gridSize {
		log.Fatalf("not enought screen space to create the grid")
	}

	for !rl.WindowShouldClose() {
		UpdateMainMenu()

		rl.BeginDrawing()

		switch screen {
		case MainMenu:
			DrawMainMenu()
		case Level1:
			DrawGrid(gridSize, gridCells, cellSize)
		}

		rl.EndDrawing()
	}
}

func DrawGrid(gridSize int, gridCells int, cellSize int) {
	rl.ClearBackground(rl.White)
	rl.DrawText("Hello", 0, 0, 24, rl.Red)

	for i := range gridCells {
		for j := range gridCells {
			rl.DrawRectangleLines(int32(screenWidth/2-gridSize/2+(i*cellSize)), int32(screenHeight/2-gridSize/2+(j*cellSize)), int32(cellSize), int32(cellSize), rl.Red)
		}
	}
}

func UpdateMainMenu() {

}

func DrawMainMenu() {
	title := func() Text {
		content := "Untitled Game"
		fontSize := int32(32)
		width := rl.MeasureText(content, fontSize)

		posX, posY := screenWidth/2-width/2, screenHeight/2-fontSize/2

		return Text{
			Content:  content,
			FontSize: fontSize,
			PosX:     posX,
			PosY:     posY,
		}
	}()

	rl.ClearBackground(rl.White)
	rl.DrawText(title.Content, title.PosX, title.PosY, title.FontSize, rl.Red)

	text := func() Text {
		content := "Click to START"
		fontSize := int32(16)
		width := rl.MeasureText(content, fontSize)
		posX := screenWidth/2 - width/2
		posY := screenHeight/2 - fontSize/2 + title.FontSize + 5

		return Text{
			Content:  content,
			FontSize: fontSize,
			PosX:     posX,
			PosY:     posY,
		}
	}()

	button := func() rl.Rectangle {
		paddingInline := int32(12)
		paddingBlock := int32(8)
		posX := float32(text.PosX - paddingInline)
		posY := float32(text.PosY - paddingBlock)

		width := float32(rl.MeasureText(text.Content, text.FontSize) + (paddingInline * 2))
		height := float32(text.FontSize + (paddingBlock * 2))

		button := rl.Rectangle{
			X:      posX,
			Y:      posY,
			Width:  width,
			Height: height,
		}

		return button
	}()

	rl.DrawRectangleRec(button, rl.Red)
	rl.DrawText(text.Content, text.PosX, text.PosY, text.FontSize, rl.Black)
}

type Text struct {
	Content  string
	FontSize int32
	PosX     int32
	PosY     int32
}
