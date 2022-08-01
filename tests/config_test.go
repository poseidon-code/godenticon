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