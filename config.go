package godenticon

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

// Read, Check & Set configurations from a JSON config file.
// Sets both Identicon.IdenticonOptions & Identicon.ImageOptions
// requires: absolute/relative path of the JSON file.
func (i *Identicon) ReadConfiguration(path string) {
    f, err := os.Open(path)
    if err!=nil {
        f.Close()
        log.Fatalln("Invalid config file path :", path)
    }
    if f.Name()[len(f.Name())-5:]!=".json" {
        log.Fatalln("Inavalid config file :", f.Name(), "(required a .json file)")
    }
    b, _ := ioutil.ReadAll(f)
    f.Close()

    identicon_o := IdenticonDefaultOptions
    json.Unmarshal(b, &identicon_o)
    identicon_o.CheckConfiguration()
    
    image_o := ImageDefaultOptions
    json.Unmarshal(b, &image_o)
    image_o.CheckConfiguration()

    i.IdenticonOptions = identicon_o
    i.ImageOptions = image_o
}

// Checks Identicon.IdenticonOptions for errors
func (o *IdenticonConfiguration) CheckConfiguration() {
    if o.Size<4 || o.Size>8 {
        log.Fatalln(
            "Invalid identicon size :", o.Size,
            "\nSize must lie between 4 to 8 (inclusive)",
        )
    }
}

// Checks Identicon.ImageOptions for errors
func (o *ImageConfiguration) CheckConfiguration() {
    if o.Size!="S" && o.Size!="M" && o.Size!="L" && o.Size!="X" {
        log.Fatalln(
            "Invalid image size :", o.Size,
            "\nImage size value (string) should be any one of S, M, L & X",
            "\ni.e.: Identicon.ImageOptions.Size='X'",
        )
    }

    r, _ := regexp.Compile(`^[a-fA-F0-9]{6}$`)
    fg_ok := r.MatchString(o.FG)
    bg_ok := r.MatchString(o.BG)
    
    if !fg_ok || !bg_ok {
        if !fg_ok { log.Println("Invalid foreground color :", o.FG) }
        if !bg_ok { log.Println("Invalid background color :", o.BG) }
        log.Fatalln(
            "Colors must be in HEX format string of length 6 (range: '000000' to 'ffffff')",
            "\ne.g.: 'ff0044'(correct) | 'f04'(wrong) | 'ff55aa00'(wrong)",
        )
    }
}

// Checks both Identicon.IdenticonOptions & Identicon.ImageOptions
// altogether for errors (kind of an One-Time function call).
// 
// When both Identicon.IdenticonOptions & Identicon.ImageOptions
// were already set by any means, then use this function directly
// instead of seperately calling Identicon.IdenticonOptions.CheckConfiguration()
// and Identicon.ImageOptions.CheckConfiguration(), to check for errors.
// 
// i.e.: Identicon.CheckConfiguration()
func (i *Identicon) CheckConfiguration() {
    i.IdenticonOptions.CheckConfiguration()
    i.ImageOptions.CheckConfiguration()
}

// Sets Identicon.IdenticonOptions & Identicon.ImageOptions
// using default values specified as variables -
// IdenticonDefaultOptions & ImageDefaultOptions for 
// identicon & image respectively.
func (i *Identicon) UseDefaultConfiguration() {
    i.IdenticonOptions  = IdenticonDefaultOptions
    i.ImageOptions      = ImageDefaultOptions
}