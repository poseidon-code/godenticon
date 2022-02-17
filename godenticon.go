package godenticon

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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

    // Creates and saves an identicon as PNG image. Requires a `path` variable
    // to be passed, as an image saving directory or name. Size of the image
    // identicon depends on the Identicon.ImageOptions.Size property which is minimum
    // of either width & height of the image, divided by some value.
    SaveImage(path string)

    // Creates and saves an identicon as SVG. Requires a `path` variable
    // to be passed, as a SVG saving directory or name. Size of the SVG
    // identicon depends on the Cell Size of every block in the identicon,
    // which is currently hardcoded to 100 units.
    SaveSVG(path string)
}


// Creates and sanitizes the saving path. Creates directory if necessary.
// path : the entire path passed as a string,
// dt   : default text to be appended always,
// ext  : extension of the file, including the '.' (i.e.: ext = ".svg" | ".png")
func handle_save_path(path, dt, ext string) (save string) {
    d, n := filepath.Split(path)

    r, _ := regexp.Compile("[^a-zA-Z0-9_-]+")
    n = r.ReplaceAllString(n, "")
    if n=="" {
        n = dt + "." + ext
    } else {
        n = n[:len(n)-3] + "." + ext
    }


    if d=="" {
        // when path is empty string, use default directory
        h, err := os.UserHomeDir()
        if err!=nil {
            fmt.Println(err)
            os.Exit(1)
        }

        d = filepath.Join(h, "Pictures")

        _, err = os.Stat(d)
        if err!=nil {
            if os.IsNotExist(err) {
                os.MkdirAll(d, os.ModePerm)
            } else {
                os.Exit(1)
            }
        }

        save = filepath.Join(d, n)
    } else {
        // when directory doesn't exists OR current working directory
        _, err := os.Stat(d)
        if err!=nil {
            if os.IsNotExist(err) {
                os.MkdirAll(d, os.ModePerm)
            } else {
                os.Exit(1)
            }
        }

        save = filepath.Join(d, n)
    }

    return save
}