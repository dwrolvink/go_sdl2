package main
// =====================================================================
// 				Imports
// =====================================================================
// Import built-in packages
import (
	"fmt"        // used for outputting to the terminal
	"time"       // used for pausing, measuring duration, etc
	"math/rand"  // random number generator
)

// Import external packages
// https://github.com/veandco/go-sdl2
// https://markkeeley.xyz/2016/go-sdl2-lesson-1/
import "github.com/veandco/go-sdl2/sdl"


// subpackages
import (
	"go_sdl2/graphicsx"
	"go_sdl2/world"
)

func main() {
	// Load SDL2, and get window and screen
	graphics := graphicsx.Initialize_graphics()
	var renderer = graphics.Renderer
	var window = graphics.Window

	// Create grid of cells (just the rects for now)
	var rect_grid = world.CreateRectGrid()

	// Not sure if this line is strictly needed
	var event sdl.Event

	// Draw loop
	var running = true
	for running	{

		/* Clear the entire screen to white. */
		// set color white
		// red, green, blue, alpha (alpha determines opaque-ness - usually 255)
		renderer.SetDrawColor(255, 255, 255, 255)
		// clear the window with specified color - in this case white.
		renderer.Clear()

		// Draw squares randomly from grid
		var r_col = rand.Intn(64)
		var r_row = rand.Intn(48)
		//screenSurface.FillRect(&rect_grid[r_row][r_col], sdl.MapRGB(screenSurface.Format, 0, 0, 0))
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.FillRect(&rect_grid[r_row][r_col])

		// Draw Screen
		//window.UpdateSurface()
			
		// The rects have been drawn, now it is time to tell the renderer to show
		// what has been draw to the screen
		renderer.Present()


		// Lag
		time.Sleep(time.Millisecond * 1)

		// handle events, in this case escape key and close window
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			// event that is sent when the window is closed
			case *sdl.QuitEvent:
				running = false
			// keydown/keyup events
			case *sdl.KeyboardEvent:
				fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
					t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
			}
		}		
	}
	
	renderer.Destroy()
	window.Destroy()

	// program is over, time to start shutting down. Keep in mind that sdl is written in C and does not have convenient
	// garbage collection like Go does
	
	sdl.Quit()
}