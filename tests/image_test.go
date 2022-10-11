package godenticon

import (
	"fmt"
	"testing"

	g "github.com/poseidon-code/godenticon"
)


type testIdenticonConfiguration struct {
    Size        []int
    Square      []bool
    Border      []bool
    Vertical    []bool
    Invert      []bool
    Symmetric   []bool
}
var tIC = testIdenticonConfiguration{
    Size:       []int{4,5,6,7,8},
    Square:     []bool{true, false},
    Border:     []bool{true, false},
    Vertical:   []bool{true, false},
    Invert:     []bool{true, false},
    Symmetric:  []bool{true, false},
}


type testImageConfiguration struct {
    Size        []string
    Portrait    []bool
    FG          []string
    BG          []string
}
var tImC = testImageConfiguration{
    Size:       []string{"S", "M", "L", "X"},
    Portrait:   []bool{true, false},
    FG:         []string{"03fcba"},
    BG:         []string{"013225"},
}

// Test for generation of images for all possible configurations
// of identicon options and image options.
// All possible combinations includes 1280 files.
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