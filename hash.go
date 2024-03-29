package godenticon

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"log"
)

// generates a hash with 1:1(square) aspect ratio
func (i *Identicon) square_hashing() {
    i.Hash      = fmt.Sprintf("%x", sha256.Sum256([]byte(i.Text)))
    i.width     = 1
    i.height    = 1
}

// generates a hash with 2:1(wide) aspect ratio
func (i *Identicon) wide_hashing() {
    i.Hash      = fmt.Sprintf("%x", sha512.Sum512([]byte(i.Text)))
    i.width     = 2
    i.height    = 1
}

// Sets Identicon.Hash, Indenticon.Width & Identicon.Height
// based on Identicon.IdenticonOptions.Square (bool).
// 
// if: Identicon.IdenticonOptions.Square == TRUE, then
// generate(set) Identicon.Hash suitable for square (1:1) aspect ratios
// and set Identicon.width, Identicon.height = 1, 1
// 
// else:
// generate(set) Identicon.Hash suitable for wide (2:1) aspect ratios
// and set Identicon.width, Identicon.height = 2, 1
func (i *Identicon) GenerateHash() {
    // fail-safe: if Identicon.Text is an empty string, then exit the program.
    if len(i.Text)==0 {
        log.Fatalln("Text to make hash from, is empty. \nSet Identicon.Text before calling GenerateHash().")
    }

    if i.IdenticonOptions.Square {
        i.square_hashing()
    } else {
        i.wide_hashing()
        if i.IdenticonOptions.Vertical {
            // width & height are switched, hence creating a tall identicon matrix
            // only visible for wide identicons (i.e.: aspect ratio = 1:2)
            // as there is no difference for square identicons (i.e.: aspect ratio = 1:1)
            i.width, i.height = i.height, i.width
        }
    }
}


// Checks the already set Identicon.Hash and
// sets Indenticon.Width & Identicon.Height
// based on Identicon.IdenticonOptions.Square (bool).
// 
// if: Identicon.IdenticonOptions.Square == TRUE, then
// for square (1:1) aspect ratios
// set Identicon.width, Identicon.height = 1, 1
// 
// else:
// for wide (2:1) aspect ratios
// set Identicon.width, Identicon.height = 2, 1
func (i *Identicon) CheckHash() {
    if len(i.Hash)==64 {
        if !i.IdenticonOptions.Square {
            log.Fatalf("Cannot use hash of length (%d) for generating non-square identicon \nHash must be of 128 characters long\n", len(i.Hash))
        }
        i.width = 1
        i.height = 1
    } else if len(i.Hash)==128 {
        if i.IdenticonOptions.Square {
            log.Fatalf("Cannot use hash of length (%d) for generating square identicon \nHash must be of 64 characters long\n", len(i.Hash))
        }
        i.width = 2
        i.height = 1
        if i.IdenticonOptions.Vertical {
            // width & height are switched, hence creating a tall identicon matrix
            // only visible for wide identicons (i.e.: aspect ratio = 1:2)
            // as there is no difference for square identicons (i.e.: aspect ratio = 1:1)
            i.width, i.height = i.height, i.width
        }
    } else {
        log.Fatalf("Invalid hash length (%d) \nHash must be of 64 or 128 characters long\n", len(i.Hash))
    }
}
