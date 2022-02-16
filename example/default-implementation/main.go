package main

import gi "github.com/poseidon-code/godenticon"

func main() {
    var i gi.Identicon
    
    i.UseDefaultConfiguration()
    i.Text = "defaults"

    i.GenerateHash()
    i.GenerateMatrix()
    i.Print()
    i.SaveImage("./example.png")
    i.SaveSVG("./example.svg")
}