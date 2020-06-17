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
import "github.com/veandco/go-sdl2/sdl"

// subpackages
import (
	"go_sdl2/graphicsx"
	"go_sdl2/world"
)

// This is the entry point for our app. Code execution starts here.

func main() {

	// ========= Init step =========

	// Load SDL2, and get window and renderer.
	// See the file graphicsx/graphicsx.go for more information on the
	// graphics struct, and the initialization steps.

	// Endpoint is that we have a window object that we write to (and can
	// close). And a renderer object, which does the writing.
	graphics := graphicsx.Initialize_graphics()

	var renderer = graphics.Renderer
	var window = graphics.Window

	// Load images into memory
	graphics.LoadImage("src/images/icon.png") // --> graphics.Images[0]
	graphics.LoadImage("src/images/cat.png")  // --> graphics.Images[1]

	var label_icon = graphics.Images[0]
	var cat_icon = graphics.Images[1]

	// Get screen width so that we can position images against the right
	// side. (You can change _ to screenHeight, if you want to use that too)
	screenWidth, _ := window.GetSize()

	// Create grid of rectangles. These will be drawn at random in black
	// in a later step.
	var rect_grid = world.CreateRectGrid()

	// Define variables outside of loop, so that we don't have to recreate
	// them every iteration.
	var event sdl.Event
	var r_col int
	var r_row int

	// Define variables outside of loop that we want to increment/decrement
	// every iteration
	var angle = 0.0


	// ========= Game loop =========

	// This variable allows us to exit the otherwise endless loop when we want
	var running = true

	// Endless loop unless running == false
	// One iteration of this loop is one draw cycle.
	for running	{

		// Increment angle, and loop back when it makes a full circle
		angle++
		if (angle >= 360.0){
			angle = 0.0
		}

		// set draw color to white
		renderer.SetDrawColor(255, 255, 255, 255)                        // red, green, blue, alpha (alpha = transparency)

		// clear the window with specified color - in this case white.
		renderer.Clear()

		// now set it to black so that we can draw black rectangles
		renderer.SetDrawColor(0, 0, 0, 255)

		// Draw squares randomly from premade rectangle grid
		// See world/world.go for the code to make a rectangle
		r_col = rand.Intn(64)
		r_row = rand.Intn(48)
		renderer.FillRect(&rect_grid[r_row][r_col])

		// Draw little red icon at topright
		renderer.Copy(label_icon.Texture, nil, &sdl.Rect{screenWidth - label_icon.Width, 0, label_icon.Width, label_icon.Height})

		// Draw rotating cat
		// A different way of drawing onto the screen with more options.
		// The first 3 parameters are the same as before. 
		// 4: angle in degrees. 5: point which the image rotates around 
		// 6: sdl.FLIP_NONE, sdl.FLIP_HORIZONTAL, sdl.SDL_FLIP_VERTICAL
		// Want to combine flips? Use, for example: sdl.FLIP_HORIZONTAL | sdl.SDL_FLIP_VERTICAL
		renderer.CopyEx(cat_icon.Texture, nil, &sdl.Rect{300, 200, cat_icon.Width, cat_icon.Height}, angle, nil, sdl.FLIP_HORIZONTAL)

		// Draw Screen
		// The rects have been drawn, now it is time to tell the renderer to show
		// what has been draw to the screen. "Present them."
		renderer.Present()

		// Sleep a little so that we go the speed that we want
		time.Sleep(time.Millisecond * 3)

		// Handle events, in this case escape key and close window
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
				
				// event that is sent when the window is closed
				case *sdl.QuitEvent:
					// setting running to false will end the game loop
					running = false

				// keydown/keyup events
				case *sdl.KeyboardEvent:
					// print keyevent information, and whatever you want to debug
					fmt.Printf("[%d ms] screen_width:%d Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
						t.Timestamp, screenWidth, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
			}
		}		
	} 
	
	// ========= End of Game loop =========

	// program is over, time to start shutting down. Keep in mind that sdl is written in C and does not have convenient
	// garbage collection like Go does

	renderer.Destroy()
	window.Destroy()
	sdl.Quit()
}