// This example is a pipeline like implementation for generating identicons
// programatically using just the Default configurations. All the steps are
// arranged specifically, numbered intentionally as it is and are in ordered
// for the program to work correctly. The `./example/default-implementation/example.png`
// & `./example/default-implementation/example.svg` are both generated from this
// example program.
//
// This example program works (and only compatible) with Go 1.18 and above
// due to the usage of Go workspaces. Open the root project and run like :
// go run ./example/default-implementation/main.go
//
// To run this example independently:
// 1. Copy this example file to some directory
// 2. Create go.mod file        : go mod init project_name
// 3. Install dependencies      : go get github.com/poseidon-code/godenticon
// 4. Run the program           : go run ./main.go

package main

import g "github.com/poseidon-code/godenticon"

func main() {
    // 1. Create an Identicon instance
    var identicon g.Identicon

    // 2. use (set) default configuration
    // identicon.IdenticonOptions = g.IdenticonDefaultOptions
    // identicon.ImageOptions = g.ImageDefaultOptions

    // 2(OR). use (call & set) default configuration
    identicon.UseDefaultConfiguration()

    // 3. Set the Text to be made into an identicon
    identicon.Text = "godenticon"

    // 4. Use Identicon.New() to perform all the steps in the pipeline
    identicon.New().SaveImage("./example/default-implementation/example.png")
    identicon.SaveSVG("./example/default-implementation/example.svg")
}