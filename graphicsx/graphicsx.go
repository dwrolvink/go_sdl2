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
import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/img"
)

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
	Images []Image
}

// Add a function to the struct. We'll be able to call it like so:
//		var graph = Graphics{window, screenSurface}
//		graph.ClearScreen()

func (this *Graphics) LoadImage(path string) {  
	// Create surface, this is needed to create the optimized **texture**
    surfaceImg, err := img.Load(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
		os.Exit(4)
	}	

	// Make Image struct to save the image info in
	var image = Image{}

	// This is for getting the Width and Height of surfaceImg. Once surfaceImg.Free() is called we lose the
	// ability to get information about the image we loaded into ram
	image.Width = surfaceImg.W
	image.Height = surfaceImg.H	

	// Take the surfaceImg and use it to create a hardware accelerated textureImg. Or in other words take the image
	// sitting in ram and put it onto the graphics card.
	textureImg, err := this.Renderer.CreateTextureFromSurface(surfaceImg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		os.Exit(5)
	}
	// We have the image now as a texture so we no longer have need for surface. Time to let it go
	surfaceImg.Free()	

	// save to struct
	image.Texture = textureImg

	// add image to graphics.Images
	this.Images = append(this.Images, image)
}

// =====================================================================
// 				Struct: Image
// =====================================================================

type Image struct {
	Texture *sdl.Texture
	Width int32
	Height int32
}




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

func Initialize_graphics() Graphics {
	// try to initialize everything
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize sdl: %s\n", err)
		os.Exit(1)
	}

	// Get config (for screentitle)
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

	// Output
	var graphics = Graphics{window, renderer, []Image{}}	

	// SUGGEST to sdl that it use a certain scaling quality for images. Default is "0" a.k.a. nearest pixel sampling
	// try out settings 0, 1, 2 to see the differences with the rotating stick figure. Change the
	// time.Sleep(time.Millisecond * 10) into time.Sleep(time.Millisecond * 100) to slow down the speed of the rotating
	// stick figure and get a good look at how blocky the stick figure is at RENDER_SCALE_QUALITY 0 versus 1 or 2
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")	

	// Return output
	return graphics

	
}