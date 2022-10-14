package godenticon_test

import (
	"fmt"
	"testing"

	g "github.com/poseidon-code/godenticon"
)

// Test for generation of SVGs for all possible configurations
// of identicon options (image options are not required).
// All possible combinations includes 160 files.
//
// NOTE : tIC declared in godenticon_test.go (required to be included
// when running : go test -v ./tests/svg_test.go ./tests/godenticon_test.go)
func TestSaveSVG(t *testing.T) {
    var i g.Identicon
    var identiconOptions g.IdenticonConfiguration
    var imageOptions = g.ImageDefaultOptions
    i.Text = "godenticon"
    count := 0

    // iterating over all possible identicon configurations
    for _, size         := range tIC.Size {
    for _, square       := range tIC.Square {
    for _, border       := range tIC.Border {
    for _, vertical     := range tIC.Vertical {
    for _, invert       := range tIC.Invert {
    for _, symmetric    := range tIC.Symmetric {
        // generating svgs
        identiconOptions = g.IdenticonConfiguration{
            Size        : size,
            Square      : square,
            Border      : border,
            Vertical    : vertical,
            Invert      : invert,
            Symmetric   : symmetric,
        }

        i.IdenticonOptions = identiconOptions
        i.ImageOptions = imageOptions
        i.CheckConfiguration()

        // filename properties are '-' seperated
        // i.e. Text-Size-Square-Border-Vertical-Invert-Symmetric
        // e.g.: godenticon-4-true-false-true-false-false
        save := "./svg_test/" + fmt.Sprintf("%v-%v-%v-%v-%v-%v-%v", i.Text,
            i.IdenticonOptions.Size, i.IdenticonOptions.Square, i.IdenticonOptions.Border, i.IdenticonOptions.Vertical, i.IdenticonOptions.Invert, i.IdenticonOptions.Symmetric,
        )

        i.GenerateHash()
        i.GenerateMatrix()
        i.SaveSVG(save)
        count++
        t.Logf("\n%v\t: %v\n\n", count, save)
    }
    }
    }
    }
    }
    }
}