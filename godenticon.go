package godenticon

import (
	"os"
	"path/filepath"
	"strings"
)

type Identicon struct {
    IdenticonOptions        IdenticonConfiguration
    ImageOptions            ImageConfiguration
    Text                    string
    Hash                    string
    Width, Height           int
    Matrix                  [][]int
}

type IdenticonConfiguration struct {
    Size        int         `json:"size"`               // sets size of the identicon (range: 4-8)
    Square      bool        `json:"square"`             // creates a square identicon
    Border      bool        `json:"border"`             // adds a border to the identicon
    Vertical    bool        `json:"vertical"`           // creates identicon in portrait dimension (not visible on using --square flag)
    Invert      bool        `json:"invert"`             // inverts the cell filling of identicon
    Symmetric   bool        `json:"symmetric"`          // creates symmetric identicon
}

type ImageConfiguration struct {
    Size        string      `json:"image-size"`         // saves image with given resolution preset (S,M,L,X)
    Save        bool        `json:"save"`               // save the identicon as an image with default image options
    SaveDir     string      `json:"save-dir"`           // saves image to the specified directory
    Portrait    bool        `json:"image-portrait"`     // saves image with portrait dimensions
    FG          string      `json:"fg"`                 // sets image's foreground color
    BG          string      `json:"bg"`                 // sets image's background color
}


type IdenticonIF interface {
    // Sets Identicon.IdenticonOptions & Identicon.ImageOptions
    // using default values specified as variables -
    // IdenticonDefaultOptions & ImageDefaultOptions for 
    // identicon & image respectively.
    UseDefaultConfiguration()

    // Read, Check & Set configurations from a JSON config file.
    // Sets both Identicon.IdenticonOptions & Identicon.ImageOptions.
    // (requires: absolute/relative path of the JSON file)
    ReadConfiguration(path string)

    // Check for errors in configurations.
    // Applies to IdenticonConfiguration, ImageConfiguration
    // and Identicon types.
    CheckConfiguration()

    // Sets Identicon.Hash, Indenticon.Width & Identicon.Height
    // based on Identicon.IdenticonOptions.Square (bool).
    // 
    // if: Identicon.IdenticonOptions.Square == TRUE, then
    // generate(set) Identicon.Hash suitable for square (1:1) aspect ratios
    // and set Identicon.Width, Identicon.Height = 1, 1
    // else:
    // generate(set) Identicon.Hash suitable for wide (2:1) aspect ratios
    // and set Identicon.Width, Identicon.Height = 2, 1
    GenerateHash()

    // Sets Identicon.Matrix based on some Identicon.IdenticonOptions.
    // Uses IdenticonConfiguration.Vertical (bool), IdenticonConfiguration.Invert (bool)
    // & IdenticonConfiguration.Symmetric (bool) to generate 8 different types of
    // matrices.
    // 
    // Combinations:
    // vertical-inverted-symmetric, vertical-inverted-asymmetric, vertical-original-symmetric
    // vertical-original-asymmetric, horizontal-inverted-symmetric, horizontal-inverted-asymmetric
    // horizontal-original-symmetric, horizontal-original-asymmetric
    GenerateMatrix()

    // Printing the Identicon.Matrix to the terminal.
    // Identicon.IdenticonOptions.Border (bool) is used to determine whether
    // to print the identicon with/without border.
    Print()

    SaveImage(path string)

    // Creates and saves an identicon as SVG. Requires a `path` variable
    // to be passed, as a SVG saving directory or name. Size of the SVG
    // identicon depends on the Cell Size of every block in the identicon,
    // which is currently hardcoded to 100 units.
    SaveSVG(path string)
}


func handle_save_path(path, t string) (save string) {
    d, n := filepath.Split(path)

    if len(d)==0 && len(n)==0 {
        // if no `path` was provided i.e.: path=""
        d = "./"
        n = t+".svg"
    }

    if len(d)==0 {
        // directory is not provided with the `path`
        // use base relative directory
        d = "./"
    }

    if len(n)==0 || (len(n)==1 && string(n[0])==".") {
        // file name is not provided with the `path`
        // OR, file name is "."
        // use `directory + Identicon.Text`
        n = t+".svg"
    }

    // file name has wrong extension
    // i.e.: file extension is anything other than `.svg`
    fn := strings.Split(n, ".")
    if len(n)!=0 && fn[len(fn)-1]!="svg" {
        // check & sanitize file extension with `.svg`
        n = strings.Join(fn[:len(fn)-1], "_")+".svg"
    }

    os.MkdirAll(d, os.ModePerm)
    save = d+n
    return save
}