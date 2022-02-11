package godenticon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Read, Check & Set configurations from a JSON config file.
// Sets booth Identicon.IdenticonOptions & Identicon.ImageOptions
// requires: absolute/relative path of the JSON file.
func (i *Identicon) ReadConfiguration(path string) {
    f, err := os.Open(path)
    if err!=nil {
        fmt.Println("Invalid config file path :", path)
        f.Close()
        os.Exit(1)
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
        fmt.Println("Invalid identicon size :", o.Size)
        fmt.Println("Size must lie between 4 to 8 (inclusive)")
        os.Exit(1)
    }
}

// Checks Identicon.ImageOptions for errors
func (o *ImageConfiguration) CheckConfiguration() {
    if o.Size!="S" && o.Size!="M" && o.Size!="L" && o.Size!="X" {
        fmt.Println("Invalid image size :", o.Size)
        fmt.Println("Image size can be one of S, M, L & X")
        os.Exit(1)
    }

    if _, err := os.Stat(o.SaveDir); err!=nil {
        if os.IsNotExist(err) {
            fmt.Println("Invalid saving directory :", o.SaveDir)
            fmt.Println("Directory doesn't exists")
            os.Exit(1)
        }
    }

    if len(o.FG)!=6 || len(o.BG)!=6 {
        if len(o.FG)!=6 {
            fmt.Println("Invalid foreground color :", o.FG)
        }
        if len(o.BG)!=6 {
            fmt.Println("Invalid background color :", o.BG)
        }
        fmt.Println("Colors must be in HEX format string of length 6 (range: '000000' to 'ffffff')")
        fmt.Println("e.g.: 'ff0044'(correct) | 'f04'(wrong) | 'ff55aa00'(wrong)")
        os.Exit(1)
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