// A general main program to test the the current implementaions of the
// godenticon package. This program works (and only compatible) with
// Go 1.18 and above due to the usage of Go workspaces.

package main

import (
	"github.com/poseidon-code/godenticon"
)

func main() {
	var i godenticon.Identicon
	i.UseDefaultConfiguration()
	i.Text = "godenticon"
	i.CheckConfiguration()
	i.GenerateHash()
	i.GenerateMatrix()
	i.SaveImage("./tests/run/run_image.png")
	i.SaveSVG("./tests/run/run_svg.svg")
	i.Print()
}