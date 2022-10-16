// This example is a pipeline like implementation for generating identicons
// programatically. All the steps are arranged specifically, numbered intentionally
// as it is and are in ordered for the program to work correctly.
// This example is an overall intendent use of this package `godenticon`.
// The `./example/example.png` & `./example/example.svg` are both generated
// from this example program.
//
// This example program works (and only compatible) with Go 1.18 and above
// due to the usage of Go workspaces. Open the root project and run like :
// go run ./example/all-implementations.go
//
// To run this example independently:
// 1. Copy this example file & `example-config.json` file to some directory
// 2. Create go.mod file        : go mod init project_name
// 3. Install dependencies      : go get github.com/poseidon-code/godenticon
// 4. Run the program           : go run ./all-implementations.go

package main

import (
	g "github.com/poseidon-code/godenticon"
)

func main() {
    // 1. Create an Identicon instance
    var identicon g.Identicon

    // 2. Create IdenticonConfiguration & ImageConfiguration
    identiconOptions := g.IdenticonConfiguration{
        Size: 8,
        Border: false,
        Square: false,
        Vertical: false,
        Invert: true,
        Symmetric: true,
    } // all fields are compulsory

    imageOptions := g.ImageConfiguration{
        Size: "X",
        Portrait: false,
        FG: "6dff24",
        BG: "0b2100",
    } // all fields are compulsory

    // 2(OR). Read, Check and Set a JSON configuration file
    // identicon.ReadConfiguration("./example-config.json")

    // 3. Check configuration (compulsory)
    // (NOTE: JSON configuration will be automatically checked)
    // identiconOptions.CheckConfiguration()
    // imageOptions.CheckConfiguration()

    // 4. Set identicon with Identicon & Image configurations
    // (NOTE: JSON configuration will be automatically set)
    identicon.IdenticonOptions = identiconOptions
    identicon.ImageOptions = imageOptions

    // 3(OR). Check entire identicon configuration if both
    // identicon.IdenticonOptions & identicon.ImageOptions
    // are already set
    identicon.CheckConfiguration()
    
    // 4(OR). use (set) default configuration
    // identicon.IdenticonOptions = g.IdenticonDefaultOptions
    // identicon.ImageOptions = g.ImageDefaultOptions
    
    // 4(OR). use (call & set) default configuration
    // identicon.UseDefaultConfiguration()

    // 5. Set the Text to be made into an identicon
    identicon.Text = "godenticon"

    // 5(OR). Set the Hash to be made into identicon
    // (NOTE: Hash must be of 64/128 characters long)
    // identicon.Hash = "2b0ff8c83f4a9c83f4a921d36ce9ce47d0d13c5d85f21d36ce9ce47d0d13c5d8"

    // 6. Generate Hash which sets identicon.Hash and both the 
    // identicon.width & identicon.height aspect ratio values
    identicon.GenerateHash()

    // 6(OR). Check Hash if manually setting the identicon.Hash
    // identicon.CheckHash()

    // 7. Generate Matrix which sets identicon.Matrix
    identicon.GenerateMatrix()

    // 8. Print the matrix to the terminal
    identicon.Print()

    // 9. Save an Image of that identicon
    identicon.SaveImage("./example/example.png")

    // 9. Save a SVG of that identicon
    identicon.SaveSVG("./example/example.svg")
}