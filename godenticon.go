package godenticon

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type Identicon struct {
    IdenticonOptions        IdenticonConfiguration
    ImageOptions            ImageConfiguration
    Text                    string
    Hash                    string
    width, height           int
    matrix                  [][]int
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
    // and set Identicon.width, Identicon.height = 1, 1
    // else:
    // generate(set) Identicon.Hash suitable for wide (2:1) aspect ratios
    // and set Identicon.width, Identicon.height = 2, 1
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


// Handles all the intermediate steps for generating an identicon from 
// Identicon configuration. It requires to have the Identicon configuration
// (Identicon, ImageConfiguration, IdenticonConfiguration) options to be set.
// Typically used for programmatically generating identicons.
// Chainable to allow New().SaveImage() & New().SaveSVG().
func (i *Identicon) New() *Identicon {
    i.CheckConfiguration()
    if i.Hash=="" {
        i.GenerateHash()
    } else {
        i.CheckHash()
    }
    i.GenerateMatrix()
    i.Print()

    return i
}


// Creates and sanitizes the saving path. Creates directory if necessary.
// path : the entire path passed as a string,
// ext  : extension of the file
func handleSavePath(path, ext string) (save string) {
	d, n := filepath.Split(path)
	r, _ := regexp.Compile("[^a-zA-Z0-9_-]+")
	
	// remove extension & sanitise file name
	n = strings.Split(n,".")[0]
    n = r.ReplaceAllString(n, "")

	// set default save file name
    if n=="" {n = fmt.Sprintf("%x", time.Now().UTC().UnixNano())}

	// set default save directory
	if d=="" {
		h, err := os.UserHomeDir()
		if err!=nil {log.Fatalln(err)}
		d = filepath.Join(h, "Pictures")
	}

	// create directory if passed directory doesn't exits
	_, err := os.Stat(d)
	if os.IsNotExist(err) {
		os.MkdirAll(d, os.ModePerm)
	}

	save = filepath.Join(d, n + "." + ext)
	fmt.Println("Saved to:", save)
	return save
}