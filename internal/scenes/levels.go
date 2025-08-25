package scenes

import (
	"image/color"
	"log"

	"github.com/OkaniYoshiii/raylib-go-test/internal/entities"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MaxHomes int = 5
)

type GridInventory struct {
	Empty int
	Homes int
}

type Grid struct {
	rl.Rectangle

	Color color.RGBA
	Cells []entities.Cell
}

func (grid *Grid) Inventory() GridInventory {
	empty := 0
	homes := 0

	for _, cell := range grid.Cells {
		switch cell.Type {
		case entities.Empty:
			empty++
		case entities.Home:
			homes++
		default:
			continue
		}
	}

	return GridInventory{
		Empty: empty,
		Homes: homes,
	}
}

type LevelOne struct {
	Grid Grid
}

func NewLevelOne(screenWidth int, screenHeight int) LevelOne {
	grid := func() Grid {
		grid := Grid{}

		size := 0
		if screenWidth > screenHeight {
			size = int(float32(screenHeight) * 0.85)
		} else {
			size = int(float32(screenWidth) * 0.85)
		}

		// If int not divisable by 2 then it's even so we add one to fix it
		// Prevent placing grid on a subpixel
		if size%2 != 0 {
			size++
		}

		if screenWidth < size || screenHeight < size {
			log.Fatalf("not enought screen space to create the grid")
		}

		grid.X = float32(screenWidth/2 - size/2)
		grid.Y = float32(screenHeight/2 - size/2)

		grid.Width = float32(size)
		grid.Height = float32(size)

		grid.Color = rl.Red

		return grid
	}()

	cells := func() []entities.Cell {
		count := 400
		cells := make([]entities.Cell, count)
		rowCount := 20
		cellWidth := int(grid.Width) / rowCount
		cellHeight := int(grid.Height) / (count / rowCount)
		for i := range cells {
			row := i / rowCount
			col := i % rowCount

			cells[i].X = float32(int(grid.X) + col*cellWidth)
			cells[i].Y = float32(int(grid.Y) + row*cellHeight)
			cells[i].Width = float32(cellWidth)
			cells[i].Height = float32(cellHeight)
		}

		return cells
	}()

	grid.Cells = cells

	return LevelOne{
		Grid: grid,
	}
}

func (lvl *LevelOne) Update() {
	inventory := lvl.Grid.Inventory()

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		for i := range len(lvl.Grid.Cells) {
			if inventory.Homes < MaxHomes && rl.CheckCollisionPointRec(rl.GetMousePosition(), lvl.Grid.Cells[i].Rectangle) {
				lvl.Grid.Cells[i].Type = entities.Home
			}
		}
	}
}

func (lvl *LevelOne) Draw() {
	rl.ClearBackground(rl.White)

	for _, cell := range lvl.Grid.Cells {
		cell.Draw()
	}
}
