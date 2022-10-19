// A general main program to test the the current implementaions of the
// godenticon package. This program works (and only compatible) with
// Go 1.18 and above due to the usage of Go workspaces.

package main

import (
	g "github.com/poseidon-code/godenticon"
)

// Using the Identicon.New() to generate identicons.
// Typically used for programmatically generating identicons.
func useNew() {
    // required
    imageOptions := g.ImageConfiguration{
        Size: "X",          //required
        Portrait: false,
        FG: "6dff24",       //required
        BG: "0b2100",       //required
    }

    // required
    identiconOptions := g.IdenticonConfiguration{
        Size: 6,            //required
        Square: true,
        Border: true,
        Vertical: false,
        Invert: false,
        Symmetric: true,
    }

    var i = g.Identicon{
        IdenticonOptions: identiconOptions,     // required
        ImageOptions: imageOptions,             // required
        Text: "godenticon",                     // required (optional if Hash is provided)
        // Hash: "2b0ff8c83f4a9c83f4a921d36ce9ce47d0d13c5d85f21d36ce9ce47d0d13c5d8",    // optional
    }

    i.New().SaveImage("./tests/run/run_image.png")  // New() is chainable (generally used for SaveImage() & SaveSVG())
    i.SaveSVG("./tests/run/run_svg.svg")            // `i` contains everything required to save as image/svg
}


// Using intended pipeline like usage to generate identicons.
// Typically used to have granular control on every step of 
// generating identicons (e.g. using as a CLI like program)
func usePipeline() {
    var i g.Identicon
    i.UseDefaultConfiguration()
    i.Text = "godenticon"

    i.CheckConfiguration()
    i.GenerateHash()
    i.GenerateMatrix()
    i.Print()

    i.SaveImage("./tests/run/run_image.png")
    i.SaveSVG("./tests/run/run_svg.svg")
}


func main() {
    usePipeline()
    useNew()
}