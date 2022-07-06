package godenticon

import (
	"testing"

	g "github.com/poseidon-code/godenticon"
)

// Testing whether the string passed in Identicon.Text
// is valid string or not along with the proper hash generated
// due to Identicon.IdenticonOptions.Square property
// The square_hashing() & wide_hashing() functions are called internally
// which sets the aspect ratio of the identicon
// (which are not required to be tested as they are used internally)
func TestGenerateHash(t *testing.T) {
    texts := []string{
        // PASSED
        "input_text",
        "a",
        "/",
        "\\",
        "\"",
        "'",
        "&",

        // FAILED
        // "", // empty string
    }

    var i g.Identicon
    for _, v := range texts {
        for _, isSquare := range []bool{true, false} {
            i.Text = v
            i.IdenticonOptions.Square = isSquare
            i.GenerateHash()
            t.Log("String: ", v, "\nHash generated: ", i.Hash)
        }
    }
}
