package godenticon

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
    // using default values specified as variables in 'defaults.go' -
    // IdenticonDefaultOptions & ImageDefaultOptions for 
    // identicon & image respectively.
    UseDefaultConfiguration()

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
}


type ConfigurationIF interface {
    // Read & Set configurations from JSON file.
    // Applies to both IdenticonConfiguration & ImageConfiguration
    ReadConfiguration(path string)

    // Check for errors in configurations.
    // Applies to both IdenticonConfiguration & ImageConfiguration
    CheckConfiguration()
}