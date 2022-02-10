package main

import (
	g "github.com/poseidon-code/godenticon"
)

func main() {
    // 1. Create IdenticonConfiguration & ImageConfiguration
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

    // 1(OR). Read a JSON configuration file
    var json_idn_opts g.IdenticonConfiguration
    var json_img_opts g.ImageConfiguration
    json_idn_opts.ReadConfiguration("./example-config.json")
    json_img_opts.ReadConfiguration("./example-config.json")

    // 2. Check configuration (compulsory)
    idn_opts.CheckConfiguration()
    img_opts.CheckConfiguration()
    
    // 2(OR) Check parsed JSON configs too
    json_idn_opts.CheckConfiguration()
    json_img_opts.CheckConfiguration()

    // 3. Create an Identicon instance
    var identicon g.Identicon

    // 4. Set identicon with Identicon & Image configurations
    identicon.IdenticonOptions = idn_opts
    identicon.ImageOptions = img_opts
    
    // 4(OR). Set configurations using the read JSON config file
    identicon.IdenticonOptions = json_idn_opts
    identicon.ImageOptions = json_img_opts
    
    // 4(OR). use default configuration, i.e.:
    identicon.IdenticonOptions = g.IdenticonDefaultOptions
    identicon.ImageOptions = g.ImageDefaultOptions

    // 5. Set the Text to be made into an identicon
    identicon.Text = "poseidon"

    // 6. Generate Hash which sets identicon.Hash and both the 
    // identicon.Width & identicon.Height aspect ratio values
    identicon.GenerateHash()

    // 7. Generate Matrix which sets identicon.Matrix
    identicon.GenerateMatrix()

    // 8. Print the matrix to the terminal
    identicon.Print()
}