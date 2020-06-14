package main

//imported for logging if something went wrong
import "fmt"
import "math/rand"

//imported for logg
//imported to keep the windoing if something went wrong and to exit program if there was a problem
import "os"
import "time"

//imported to use sdl2
import "github.com/veandco/go-sdl2/sdl"


var event sdl.Event

// Made so we can initialize everything in initialize_graphics() and send this struct back
// as the result
type Graphics struct {
	window *sdl.Window
	screen *sdl.Surface
}

func (this Graphics) ClearScreen() {  
    this.screen.FillRect(nil, sdl.MapRGB(this.screen.Format, 0xff, 0xff, 0xff));
}



// Loads SDL2
// This is used to:
// - create a window
// - draw on the window
// - clear window
//
// https://github.com/veandco/go-sdl2
// https://markkeeley.xyz/2016/go-sdl2-lesson-1/

func initialize_graphics() Graphics {
	// try to initialize everything
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize sdl: %s\n", err)
		os.Exit(1)
	}

	// try to create a window
	window, err := sdl.CreateWindow("Go + SDL2 Lesson 1", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		640, 480, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(2)
	}

	// window has been created, now need to get the window surface to draw on window
	screenSurface, err := window.GetSurface()
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to create surface: %s\n", err)
		os.Exit(2)
	}

	return Graphics{window, screenSurface}
}

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



func main() {
	// Load SDL2 
	// https://github.com/veandco/go-sdl2
	// https://markkeeley.xyz/2016/go-sdl2-lesson-1/
	graphics := initialize_graphics()
	var screenSurface = graphics.screen
	var window = graphics.window

	// Create grid of cells (just the rects for now)
	var rect_grid = CreateRectGrid()

	// Draw loop
	var running = true
	for running	{

		/* Clear the entire screen to our selected color. */
		graphics.ClearScreen()

		// Draw squares randomly from grid
		var r_col = rand.Intn(64)
		var r_row = rand.Intn(48)
		screenSurface.FillRect(&rect_grid[r_row][r_col], sdl.MapRGB(screenSurface.Format, 0, 0, 0))

		// Draw Screen
		window.UpdateSurface()

		// Lag
		time.Sleep(time.Millisecond * 1)

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
					t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
			}
		}		
	}
	
	window.Destroy()

	// program is over, time to start shutting down. Keep in mind that sdl is written in C and does not have convenient
	// garbage collection like Go does
	
	sdl.Quit()
}