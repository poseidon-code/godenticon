package godenticon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Sets Identicon.IdenticonOptions from JSON config file
func (o *IdenticonConfiguration) ReadConfiguration(path string) {
    f, _ := os.Open(path); defer f.Close()
    b, _ := ioutil.ReadAll(f)
    *o = IdenticonDefaultOptions
    json.Unmarshal(b, &o)
}

// Check Identicon.IdenticonOptions for error
func (o *IdenticonConfiguration) CheckConfiguration() {
    if o.Size<4 || o.Size>8 {
        fmt.Println("Invalid identicon size :", o.Size)
        fmt.Println("Size must lie between 4 to 8 (inclusive)")
        os.Exit(1)
    }
}

// Sets Identicon.ImageOptions from JSON config file
func (o *ImageConfiguration) ReadConfiguration(path string) {
    f, _ := os.Open(path); defer f.Close()
    b, _ := ioutil.ReadAll(f)
    *o = ImageDefaultOptions
    json.Unmarshal(b, &o)
}

// Check Identicon.ImageOptions for error
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

func (i *Identicon) UseDefaultConfiguration() {
    i.IdenticonOptions  = IdenticonDefaultOptions
    i.ImageOptions      = ImageDefaultOptions
}