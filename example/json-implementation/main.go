package main

import gi "github.com/poseidon-code/godenticon"

func main() {
    var i gi.Identicon

    i.ReadConfiguration("./example-config.json")

    i.Text = "json"

    i.GenerateHash()
    i.GenerateMatrix()
    i.Print()
    i.SaveImage("./example.png")
    i.SaveSVG("./eaxmple.svg")
}