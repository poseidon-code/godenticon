package godenticon_test

import (
	"fmt"
	"testing"

	g "github.com/poseidon-code/godenticon"
)

// Test for generation of images for all possible configurations
// of identicon options and image options.
// All possible combinations includes 1280 files.
//
// NOTE : tIC & tImC declared in godenticon_test.go (required to be included
// when running : go test -v ./tests/image_test.go ./tests/godenticon_test.go)
func TestSaveImage(t *testing.T) {
    var i g.Identicon
    var identiconOptions g.IdenticonConfiguration
    var imageOptions g.ImageConfiguration
    i.Text = "godenticon"
    count := 0


    // iterating over all possible identicon configurations
    for _, size         := range tIC.Size {
    for _, square       := range tIC.Square {
    for _, border       := range tIC.Border {
    for _, vertical     := range tIC.Vertical {
    for _, invert       := range tIC.Invert {
    for _, symmetric    := range tIC.Symmetric {
    // iterating over all possible image configuartions
    for _, image_size   := range tImC.Size {
    for _, portrait     := range tImC.Portrait {
    for _, fg           := range tImC.FG {
    for _, bg           := range tImC.BG {
        // generating images
        identiconOptions = g.IdenticonConfiguration{
            Size        : size,
            Square      : square,
            Border      : border,
            Vertical    : vertical,
            Invert      : invert,
            Symmetric   : symmetric,
        }
        imageOptions = g.ImageConfiguration{
            Size        : image_size,
            Portrait    : portrait,
            FG          : fg,
            BG          : bg,
        }

        i.ImageOptions = imageOptions
        i.IdenticonOptions = identiconOptions
        i.CheckConfiguration()

        // filename properties are '-' seperated
        // i.e. Text-Size-Square-Border-Vertical-Invert-Symmetric-ImageSize-Portrait-FG-BG
        // e.g.: godenticon-4-true-false-true-false-false-L-false-03fcba-013225
        save := "./image_test/" + fmt.Sprintf("%v-%v-%v-%v-%v-%v-%v-%v-%v-%v-%v", i.Text,
            i.IdenticonOptions.Size, i.IdenticonOptions.Square, i.IdenticonOptions.Border, i.IdenticonOptions.Vertical, i.IdenticonOptions.Invert, i.IdenticonOptions.Symmetric,
            i.ImageOptions.Size, i.ImageOptions.Portrait, i.ImageOptions.FG, i.ImageOptions.BG,
        )

        i.GenerateHash()
        i.GenerateMatrix()
        i.SaveImage(save)
        count++
        t.Logf("\n%v\t: %v\n\n", count, save)
    }
    }
    }
    }
    }
    }
    }
    }
    }
    }
}