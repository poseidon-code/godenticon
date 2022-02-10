package godenticon

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

// generates a hash with 1:1(square) aspect ratio
func (i *Identicon) square_hashing() {
    i.Hash      = fmt.Sprintf("%x", sha256.Sum256([]byte(i.Text)))
    i.Width     = 1
    i.Height    = 1
}

// generates a hash with 2:1(wide) aspect ratio
func (i *Identicon) wide_hashing() {
    i.Hash      = fmt.Sprintf("%x", sha512.Sum512([]byte(i.Text)))
    i.Width     = 2
    i.Height    = 1
}

func (i *Identicon) GenerateHash() {
    if i.IdenticonOptions.Square {
        i.square_hashing()
    } else {
        i.wide_hashing()
        if i.IdenticonOptions.Vertical {
            // width & height are switched, hence creating a tall identicon matrix
            // only visible for wide identicons (i.e.: aspect ratio = 1:2)
            // as there is no difference for square identicons (i.e.: aspect ratio = 1:1)
            i.Width, i.Height = i.Height, i.Width
        }
    }
}