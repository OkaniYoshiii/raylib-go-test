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

type HomeCell struct {
	Cell
}

func (home *HomeCell) Draw() {
	color := rl.Green
	rl.DrawRectangleRec(home.Cell.Rectangle, color)
}

type EmptyCell struct {
	Cell
}

func (empty *EmptyCell) Draw() {
	color := rl.Red
	rl.DrawRectangleLines(int32(empty.Cell.X), int32(empty.Cell.Y), int32(empty.Cell.Width), int32(empty.Cell.Height), color)
}

type CellInterface interface {
	Draw()
	ChangeTo(cellType CellType)
}
