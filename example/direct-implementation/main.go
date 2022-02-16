package main

import gi "github.com/poseidon-code/godenticon"

func main() {
    var i gi.Identicon

    identicon_o := gi.IdenticonConfiguration{
        Size: 7,
        Square: false,
        Border: true,
        Vertical: true,
        Invert: false,
        Symmetric: true,
    }
    identicon_o.CheckConfiguration()

    image_o := gi.ImageConfiguration{
        Size: "L",
        Portrait: false,
        FG: "61ff26",
        BG: "091c02",
    }

    i.IdenticonOptions = identicon_o
    i.ImageOptions = image_o
    i.Text = "direct"

    i.GenerateHash()
    i.GenerateMatrix()
    i.Print()
    i.SaveImage("./example.png")
    i.SaveSVG("./example.svg")
}