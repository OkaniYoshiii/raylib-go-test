package entities

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CellType int

const (
	Empty CellType = iota
	Home
)

type Cell struct {
	rl.Rectangle

	Type CellType
}

func (cell *Cell) Draw() {
	switch cell.Type {
	case Empty:
		color := rl.Red
		rl.DrawRectangleLines(int32(cell.X), int32(cell.Y), int32(cell.Width), int32(cell.Height), color)
	case Home:
		color := rl.Green
		rl.DrawRectangleRec(cell.Rectangle, color)
	}
}
