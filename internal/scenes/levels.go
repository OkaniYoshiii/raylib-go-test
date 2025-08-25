package scenes

import (
	"image/color"

	"github.com/OkaniYoshiii/raylib-go-test/internal/entities"
	"github.com/OkaniYoshiii/raylib-go-test/internal/ui"
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

type HUD struct {
	InventoryBar ui.InventoryBar
}

type LevelOne struct {
	Grid Grid
	HUD  HUD
}

func NewLevelOne(screenWidth int, screenHeight int) LevelOne {
	inventoryBar := ui.NewInventoryBar(screenWidth, screenHeight)

	grid := func() Grid {
		grid := Grid{}

		size := 400

		grid.X = float32(screenWidth/2 - size/2)
		grid.Y = inventoryBar.Y - float32(size) - 20

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
		HUD: HUD{
			InventoryBar: inventoryBar,
		},
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

	lvl.HUD.InventoryBar.Draw()
}
