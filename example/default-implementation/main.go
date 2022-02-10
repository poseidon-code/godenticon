package main

import gi "github.com/poseidon-code/godenticon"

func main() {
    var i gi.Identicon
    
    i.UseDefaultConfiguration()
    i.Text = "default options example"

    i.GenerateHash()
    i.GenerateMatrix()
    i.Print()
}