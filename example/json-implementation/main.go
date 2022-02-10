package main

import gi "github.com/poseidon-code/godenticon"

func main() {
    var i gi.Identicon
    var o gi.IdenticonConfiguration

    o.ReadConfiguration("./example-config.json")
    o.CheckConfiguration()

    i.IdenticonOptions = o
    i.Text = "json config example"

    i.GenerateHash()
    i.GenerateMatrix()
    i.Print()
}