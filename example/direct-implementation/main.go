package main

import gi "github.com/poseidon-code/godenticon"

func main() {
    var i gi.Identicon

    o := gi.IdenticonConfiguration{
        Size: 7,
        Square: false,
        Border: true,
        Vertical: true,
        Invert: false,
        Symmetric: true,
    }
    o.CheckConfiguration()

    i.IdenticonOptions = o
    i.Text = "direct usage example"

    i.GenerateHash()
    i.GenerateMatrix()
    i.Print()
}