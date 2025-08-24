package scenes

import (
	"image/color"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Grid struct {
	rl.Rectangle

	Color       color.RGBA
	RowCount    int
	ColumnCount int
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

		grid.RowCount, grid.ColumnCount = 25, 25

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

	return LevelOne{
		Grid: grid,
	}
}

func (lvl *LevelOne) Update() {

}

func (lvl *LevelOne) Draw() {
	rl.ClearBackground(rl.White)

	cellWidth := int(lvl.Grid.Width) / lvl.Grid.RowCount
	cellHeight := int(lvl.Grid.Height) / lvl.Grid.ColumnCount

	for i := range lvl.Grid.RowCount {
		for j := range lvl.Grid.ColumnCount {
			x := int32(int(lvl.Grid.X) + (i * cellWidth))
			y := int32(int(lvl.Grid.Y) + (j * cellHeight))

			rl.DrawRectangleLines(x, y, int32(cellWidth), int32(cellHeight), lvl.Grid.Color)
		}
	}
}
