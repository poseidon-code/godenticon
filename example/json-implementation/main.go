// This example is a pipeline like implementation for generating identicons
// using JSON configuration file. All the steps are arranged specifically,
// numbered intentionally as it is and are in ordered for the program to
// work correctly. The `./example/json-implementation/example.png` &
// `./example/json-implementation/example.svg` are both generated from this
// example program.
//
// This example program works (and only compatible) with Go 1.18 and above
// due to the usage of Go workspaces. Open the root project and run like :
// go run ./example/json-implementation/main.go
//
// To run this example independently:
// 1. Copy this example file & `example-config.json` file to some directory
// 2. Create go.mod file        : go mod init project_name
// 3. Install dependencies      : go get github.com/poseidon-code/godenticon
// 4. Run the program           : go run ./main.go

package main

import g "github.com/poseidon-code/godenticon"

func main() {
    // 1. Create an Identicon instance
    var identicon g.Identicon
    
    // 2. Set the Text to be made into an identicon
    identicon.Text = "godenticon"

    // 2(OR). Set the Hash to be made into identicon
    // (NOTE: Hash must be of 64/128 characters long)
    // identicon.Hash = "2b0ff8c83f4a9c83f4a921d36ce9ce47d0d13c5d85f21d36ce9ce47d0d13c5d8"
    
    // 3. Read, Check and Set a JSON configuration file
    identicon.ReadConfiguration("./example/json-implementation/example-config.json")
    
    // 4. Generate Hash which sets identicon.Hash and both the 
    // identicon.width & identicon.height aspect ratio values
    identicon.GenerateHash()

    // 4(OR). Check Hash if manually setting the identicon.Hash
    // identicon.CheckHash()

    // 5. Generate Matrix which sets identicon.Matrix
    identicon.GenerateMatrix()

    // 6. Print the matrix to the terminal
    identicon.Print()

    // 7. Save an Image of that identicon
    identicon.SaveImage("./example/json-implementation/example.png")
    
    // 7. Save a SVG of that identicon
    identicon.SaveSVG("./example/json-implementation/example.svg")
}