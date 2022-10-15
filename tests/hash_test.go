package godenticon_test

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
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


// Testing which builtin hashing algorithms can be used,
// and whether the programs exits succesfully if hash of some
// invalid (!64 & !128) length is passed.
// The aspect ratios of the identicon are set accordingly
// (which are not required to be tested as they are used internally)
func TestCheckHash(t *testing.T) {
    var hashes = []string{
        // PASSED
        fmt.Sprintf("%x", sha256.Sum256([]byte(""))), // 64 (fails when IdenticonConfiguration.Square == false)
        fmt.Sprintf("%x", sha512.Sum512_256([]byte(""))), // 64 (fails when IdenticonConfiguration.Square == false)
        fmt.Sprintf("%x", sha512.Sum512([]byte(""))), // 128 (fails when IdenticonConfiguration.Square == true)

        // FAILED
        // fmt.Sprintf("%x", sha512.Sum384([]byte(""))), // 96 (>64 & <128)
        // fmt.Sprintf("%x", sha512.Sum512_224([]byte(""))), // 56 (<64)
        // fmt.Sprintf("%x", md5.Sum([]byte(""))), // 32 (<64)
        // fmt.Sprintf("%x", sha1.Sum([]byte(""))), // 40 (<64)
        // "c83f4a921d36ce9ce47d0d13c5d85f2b0ff8", // 36 (<64)
        // "66d8007d647c83f4a921d36ce9ceeefb8bdf1542850dd0d13c5d85f2b0ff8cf83e135720e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e", // 164 (>128)
        // "2b0ff8c83f4a9c83f4a921d36ce9ce47d0d13c5d85f21d36ce9ce47d0d13c5d85f2b0ff8", // 72 (>64 & <128)
        // "", // 0 (<64)
    }

    var i g.Identicon
    for _, v := range hashes {
        i.Hash = v
        t.Logf("Checking Hash (length: %d): %v", len(i.Hash), i.Hash)
        i.CheckHash()
    }
}
