package graphicsx
// =====================================================================
// 				Imports
// =====================================================================
// Import built-in packages
import (
	"fmt"        // used for outputting to the terminal
	"os"
)

// Import external packages
// https://markkeeley.xyz/2016/go-sdl2-lesson-1/
import "github.com/veandco/go-sdl2/sdl"

// subpackages
import (
	"go_sdl2/config"
)


// =====================================================================
// 				Struct: Graphics
// =====================================================================

// Make a struct so we can initialize everything in initialize_graphics() and send this 
// struct back as the result

type Graphics struct {
	Window *sdl.Window
	//Screen *sdl.Surface
	Renderer *sdl.Renderer
}

// Add a function to the struct. We'll be able to call it like so:
//		var graph = Graphics{window, screenSurface}
//		graph.ClearScreen()
/*
func (this Graphics) ClearScreen() {  
    this.Screen.FillRect(nil, sdl.MapRGB(this.Screen.Format, 0xff, 0xff, 0xff));
}
*/

// =====================================================================
// 				Functions
// =====================================================================

// Loads SDL2
// This function is used to:
// - initialize SDL, and handle errors
// - create a window, and get the renderer from it
// - return a struct containing the window and the renderer
//
// This struct can then be used to:
// - draw on the window
// - clear window
//
// https://github.com/veandco/go-sdl2
// https://markkeeley.xyz/2016/go-sdl2-lesson-1/

func Initialize_graphics() Graphics {
	// try to initialize everything
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize sdl: %s\n", err)
		os.Exit(1)
	}

	// Get config
	var cfg = config.GetConfig()

	// try to create a window
	window, err := sdl.CreateWindow(cfg.ScreenTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		640, 480, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(2)
	}

	// Create a renderer. 
	// A sdl.Renderer uses the hardware accelerated api available to you 
	// (DirectX, OpenGL, OpenGL ES) and gives you a cross platform way to draw 
	// graphical primitives (rectangles/lines/points) and images to the screen quickly.
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(2)
	}
	renderer.Clear()	

	// window has been created, now need to get the window surface to draw on window
	/*
		// renderer is a more efficient replacement for surface
	screenSurface, err := window.GetSurface()
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to create surface: %s\n", err)
		os.Exit(2)
	}

	return Graphics{window, screenSurface}
	*/

	return Graphics{window, renderer}

	
}