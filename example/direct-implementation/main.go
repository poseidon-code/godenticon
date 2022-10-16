// This example is a pipeline like implementation for generating identicons
// programatically. All the steps are arranged specifically,
// numbered intentionally as it is and are in ordered for the program to
// work correctly. The `./example/direct-implementation/example.png` &
// `./example/direct-implementation/example.svg` are both generated from this
// example program.
//
// This example program works (and only compatible) with Go 1.18 and above
// due to the usage of Go workspaces. Open the root project and run like :
// go run ./example/direct-implementation/main.go
//
// To run this example independently:
// 1. Copy this example file to some directory
// 2. Create go.mod file        : go mod init project_name
// 3. Install dependencies      : go get github.com/poseidon-code/godenticon
// 4. Run the program           : go run ./main.go

package main

import g "github.com/poseidon-code/godenticon"

func main() {
    // 1. Create IdenticonConfiguration & ImageConfiguration
    identiconOptions := g.IdenticonConfiguration{
        Size: 7,
        Square: false,
        Border: true,
        Vertical: true,
        Invert: false,
        Symmetric: true,
    } // all fields are compulsory

    imageOptions := g.ImageConfiguration{
        Size: "L",
        Portrait: false,
        FG: "61ff26",
        BG: "091c02",
    } // all fields are compulsory


    // 2. Create an Identicon instance
    identicon := g.Identicon{
        IdenticonOptions: identiconOptions,     // required
        ImageOptions: imageOptions,             // required
        Text: "godenticon",                     // required (optional if Hash is provided)
        // Hash: "2b0ff8c83f4a9c83f4a921d36ce9ce47d0d13c5d85f21d36ce9ce47d0d13c5d8",    // optional
    }

    // 3. Use Identicon.New() to perform all the steps in the pipeline
    identicon.New().SaveImage("./example/direct-implementation/example.png")
    identicon.SaveSVG("./example/direct-implementation/example.svg")
}