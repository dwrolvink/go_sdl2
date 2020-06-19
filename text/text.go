package text

// Import external packages
import (
	"github.com/veandco/go-sdl2/sdl"
)

// subpackages
import (
	"go_sdl2/graphicsx"
)

// =====================================================================
// 				Struct: TextObject
// =====================================================================

type TextObjectConfig struct {
	Graphics *graphicsx.Graphics            
	Text string
	Font string
	FontSize int
	Color *sdl.Color
}

func NewTextObject(toc TextObjectConfig) *TextObject {

	// Create Struct
	var textobj = TextObject{
		Graphics: toc.Graphics,             
		Text: toc.Text,
		Font: toc.Font,
		FontSize: toc.FontSize,
		Color: toc.Color,
	}		

	// Create Image
	textobj.Render()

	return &textobj
}


type TextObject struct {
	Graphics *graphicsx.Graphics
	Image *graphicsx.Image             
	Text string
	Font string
	FontSize int
	Color *sdl.Color
	Rect *sdl.Rect
	UpdateRect func(*TextObject)
}

func (this *TextObject) SetText(text string) {  
	this.Text = text
	this.Render()
}

func (this *TextObject) Render() {  
	// Create Image
	var image = this.Graphics.CreateTextImage(
						this.Text, this.Font, this.FontSize, this.Color)
	this.Image = &image
}




