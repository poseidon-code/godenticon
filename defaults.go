package godenticon

import (
	"fmt"
	"os"
)

func default_save_directory() string {
    home_dir, _ := os.UserHomeDir()
    return fmt.Sprintf("%s/Pictures", home_dir)
}

var IdenticonDefaultOptions = IdenticonConfiguration{
    Size:       6,
    Square:     false,
    Border:     false,
    Vertical:   false,
    Invert:     false,
    Symmetric:  false,
}

var ImageDefaultOptions = ImageConfiguration{
    Size:       "L",
    Save:       false,
    SaveDir:    default_save_directory(),
    Portrait:   false,
    FG:         "6dff24",
    BG:         "0b2100",
}