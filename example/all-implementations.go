package main

import (
	g "github.com/poseidon-code/godenticon"
)

func main() {
    // 1. Create an Identicon instance
    var identicon g.Identicon

    // 2. Create IdenticonConfiguration & ImageConfiguration
    idn_opts := g.IdenticonConfiguration{
        Size: 8,
        Border: false,
        Square: false,
        Vertical: false,
        Invert: true,
        Symmetric: true,
    } // all fields are compulsory

    /* Image options are not yet implemented, hence all 
     * identicon.ImageOptions related variables & assignments
     * are commented.
     */
    img_opts := g.ImageConfiguration{
        Size: "X",
        Save: true,
        SaveDir: "/path/to/directory/",
        Portrait: false,
        FG: "6dff24",
        BG: "0b2100",
    } // all fields are compulsory

    // 2(OR). Read, Check and Set a JSON configuration file
    identicon.ReadConfiguration("./example-config.json")

    // 3. Check configuration (compulsory)
    // (NOTE: JSON configuration will be automatically checked)
    idn_opts.CheckConfiguration()
    img_opts.CheckConfiguration()

    // 4. Set identicon with Identicon & Image configurations
    // (NOTE: JSON configuration will be automatically set)
    identicon.IdenticonOptions = idn_opts
    identicon.ImageOptions = img_opts
    
    // 4(OR). use default configuration, i.e.:
    identicon.IdenticonOptions = g.IdenticonDefaultOptions
    identicon.ImageOptions = g.ImageDefaultOptions

    // 5. Set the Text to be made into an identicon
    identicon.Text = "godenticon"

    // 6. Generate Hash which sets identicon.Hash and both the 
    // identicon.Width & identicon.Height aspect ratio values
    identicon.GenerateHash()

    // 7. Generate Matrix which sets identicon.Matrix
    identicon.GenerateMatrix()

    // 8. Print the matrix to the terminal
    identicon.Print()
}