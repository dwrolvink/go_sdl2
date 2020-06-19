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
import (
	"github.com/veandco/go-sdl2/sdl"
)

// subpackages
import (
	"go_sdl2/graphicsx"
	"go_sdl2/world"
	"go_sdl2/text"
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


	// Get screen dimensions so that we can position images relative to the corners.
	// Note that the screen dimensions are dictated in config/config.go, this code
	// is here to show you can get it from the window object too.
	screenWidth, screenHeight := window.GetSize()

	// Create grid of rectangles. These will be drawn at random in black
	// in a later step.
	var rect_grid = world.CreateRectGrid()

	// Define variables outside of loop, so that we don't have to recreate
	// them every iteration.
	var event sdl.Event
	var r_col int
	var r_row int

	// Text Objects
	var hello_text = text.NewTextObject(text.TextObjectConfig{
		Graphics: &graphics, 
		Text: "Kitty cat is testing your application",
		Font: "SourceCodePro-Regular.ttf", 
		FontSize: 12,
		Color: &sdl.Color{255, 0, 0, 255},
	})

	var debug_text = text.NewTextObject(text.TextObjectConfig{
		Graphics: &graphics, 
		Text: "Press a key to show keyevent",
		Font: "SourceCodePro-Regular.ttf", 
		FontSize: 12,
		Color: &sdl.Color{0, 0, 0, 120},
	})	


	// Place the hello text message
	hello_text.Rect = &sdl.Rect{
		(screenWidth - hello_text.Image.Width) / 2,      // x
		screenHeight-100,                                // y
		hello_text.Image.Width, hello_text.Image.Height,  // width, height
	}

	// Define placement function for the debug_text struct and assign it to 
	// a parameter of the struct
	debug_text.UpdateRect = func(textobj *text.TextObject) {  	
		textobj.Rect = &sdl.Rect{
			(screenWidth - textobj.Image.Width) / 2, 
			screenHeight - 80, 
			textobj.Image.Width, textobj.Image.Height,
		}
	}

	// Now we can update the Rect whenever we change the text 
	// (this will change the width, so also the Rect)
	debug_text.UpdateRect(debug_text)
	
	// Define variables outside of loop that we want to increment/decrement
	// every iteration
	var angle = 0.0  // angle of the cat image


	// ========= Game loop =========

	// This variable allows us to exit the otherwise endless loop when we want
	var running = true

	// Endless loop unless is running set to false
	// One iteration of this loop is one draw cycle.
	for running	{

		// Increment angle of the cat, and loop back when it makes a full circle
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

		// Draw squares randomly from the premade rectangle grid
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
		var kitty_rect = &sdl.Rect{(screenWidth - cat_icon.Width)/2, 60, cat_icon.Width, cat_icon.Height}
		renderer.CopyEx(cat_icon.Texture, nil, kitty_rect, angle, nil, sdl.FLIP_HORIZONTAL)

		// Draw Hello text
		renderer.Copy(hello_text.Image.Texture, nil, hello_text.Rect)

		// Draw debug text
		renderer.Copy(debug_text.Image.Texture, nil, debug_text.Rect)	
		

		// Draw Screen
		// The rects have been drawn, now it is time to tell the renderer to show
		// what has been draw to the screen. "Present them."
		renderer.Present()

		// Sleep a little so that we go the speed that we want
		time.Sleep(time.Millisecond * 3)

		// Handle events, in this case keyevents and close window
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
				
				// event that is sent when the window is closed
				case *sdl.QuitEvent:
					// setting running to false will end the game loop
					running = false

				// keydown/keyup events
				case *sdl.KeyboardEvent:
					// compile debug msg
					msg := fmt.Sprintf("[%d ms] screen_width:%d Keyboard, type:%d, sym:%c, modifiers:%d, state:%d, repeat:%d",
						t.Timestamp, screenWidth, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
				
					// show on screen
					debug_text.SetText(msg)
					debug_text.UpdateRect(debug_text)


					// print in terminal
					fmt.Println(msg)
			}
		}		
	} 
	
	// ========= End of Game loop =========

	// program is over, time to start shutting down. Keep in mind that sdl is written in C and does not have convenient
	// garbage collection like Go does
	graphics.Destroy()

}