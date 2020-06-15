package world

import "github.com/veandco/go-sdl2/sdl"

func CreateRectGrid() [48][64]sdl.Rect {
	// make color and rect grid
	//var color_grid [48][64][3]int
	var rect_grid [48][64] sdl.Rect
	//fmt.Println(grid)

	var x int32 = 0
	var y int32 = 0

	for row := 0; row < 48; row++ {
		for col := 0; col < 64; col++ {
			// color_grid[row][col][0] = rand.Intn(255)
			// color_grid[row][col][1] = rand.Intn(255)
			// color_grid[row][col][2] = rand.Intn(255)

			rect_grid[row][col] = sdl.Rect{x, y, 10, 10}
			x += 10
		}
		x = 0
		y += 10
	}

	return rect_grid
}
