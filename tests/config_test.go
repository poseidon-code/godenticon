package godenticon

import (
	"testing"

	g "github.com/poseidon-code/godenticon"
)

// Testing whether the config file path is valid/exists or not,
// if NOT then exit the program, else check whether the file type is of '.json'
// if NOT then exit the program, else set configuration options to the values
// of that of the config file.
// Invalid keys are omitted and invalid values are handled in CheckConfiguration(),
// if invalid values are found the exit program with error
func TestReadConfiguration(t *testing.T) {
	path := []string{
		// PASSED
		"valid.json",
		"./valid.json",
		
		// FAILED
		// "",					// invalid config file path
		// ".",					// invalid config file path
		// "./",				// invalid config file path
		// "invalid",			// invalid config file
		// "./invalid",			// invalid config file
		// "invalid.txt",		// invalid config file
		// "./invalid.txt",		// invalid config file
		// "invalid.json",		// invalid data, (omitted if not handled by CheckConfiguration())
		// "./invalid.json",	// invalid data, (omitted if not handled by CheckConfiguration())
	}
	var i g.Identicon

	for _, p := range path {
		i.ReadConfiguration("./config_test/"+p)
	}

	t.Log(i.IdenticonOptions)
	t.Log(i.ImageOptions)
}

// Testing the IdenticonOptions, here only the Size option is varied
// and handled, as all other options are of bool type hence does not 
// require testing.
// The IdenticonOptions.Size must lie between 4 to 8 (inclusive),
// where float like 4.0, 5.0,... etc. gets implicitly type casted to int.
func TestCheckIdenticonOptions(t *testing.T) {
	sizes := []int{
		// PASSED
		4.0,   	// implicit type casted
		4, 
		5, 
		6, 
		7,
		8, 
		8.0, 	// implicit type casted

		// FAILED
		// -1, 		// invalid (not in range 4 - 8 incl.)
		// 0, 		// invalid (not in range 4 - 8 incl.)
		// 1,  		// invalid (not in range 4 - 8 incl.)
		// 9, 		// invalid (not in range 4 - 8 incl.)
		// 9.0,		// implicit type casted, invalid (not in range 4 - 8 incl.)
	}
	
	var i g.Identicon
	o := i.IdenticonOptions
	
	for _, s := range sizes {
		o.Size = s
		o.CheckConfiguration()
	}

	t.Log(o)
}

// Testing the ImageOptions, here the valid Sizes are "S", "M", "L", "X",
// all in uppercase and everything else is invalid.
// Both the foreground (FG) & background (BG) colors are checked against 
// reular expression `^[a-fA-F0-9]{6}$` which discards any combinations
// of characters other than in the range '000000' to 'ffffff'. It accepts
// only 6 characters which can be either of lowercase or uppercase.
// The Portrait option doesn't requires checking as it is of bool type.
func TestCheckImageOptions(t *testing.T) {
	sizes := []string{
		// PASSED
		"S",
		"M",
		"L",
		"X",

		// FAILED
		// "",
		// "  ",
		// "/",
		// "\\",
		// "s",
		// "m",
		// "l",
		// "x",
		// "smlx",
		// "SMLX",
		// "a",
		// "abc",
	}

	colors := []string{
		// PASSED
		"000000",
		"FFFFFF",
		"abcdef",
		"012345",
		"678901",
		"012edf",
		
		// FAILED (invalid strings as per regex : `^[a-fA-F0-9]{6}$`)
		// "",
		// "1",
		// "22",
		// "333",
		// "4444",
		// "55555",
		// "g",
		// "      ",
		// "______",
		// "ccc",
		// "hjk",
		// "hjklio",
		// "fffffff",
	}

	var i g.Identicon
	o := i.ImageOptions

	for _, s := range sizes {
		o.BG, o.FG = "ffffff", "000000" 	// temporary colors
		o.Size = s
		o.CheckConfiguration()
	}

	for _, c := range colors {
		o.Size = "S"	// temporary size
		o.BG, o.FG = c, c
		o.CheckConfiguration()
	}

	t.Log(o)
}