package godenticon_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"
)

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

// Testing the save file path for saving the image & svg.
// 
// If no path is passed (i.e. path="") a default directory
// (HOME directory of the OS) is set and a Unix Timestamp 
// in Hexadecimal format is set as the name of the file.
// 
// If only the file name is passed (i.e. path="xyz" | "xyz.ext")
// then the default directory (HOME directory of the OS) is used.
// 
// If only the directory is passed (i.e. path="./" | "/xyz/a/b/")
// then if the directory doesn't exists, it creates the directory
// and a Unix Timestamp in Hexadecimal format is set as the file name.
// 
// If both the file nme and the directory name is passed 
// (i.e. path="./xyz.ext" | "./a/b/c/xyz.ext") then if the directory 
// doesn't exists, it creates the directory and sets the file name as passed.
func TestHandleSavePath(t *testing.T) {
	paths := []string{
		// PASSED
		"",
		".",
		"/",
		"./",
		"..",
		"../",
		"../../",
		"./p/g/",
		"./p/g.go",
		"./p/g/s.ext",
		"s.ext",

		// FAILED (only fails when user's Home directory doesn't exists)
	}

	for _, p := range paths {
		save := handleSavePath(p, "ext")
		d,n := filepath.Split(p)

		t.Logf("\nd:%s \nn:%s \nsave:%s \n\n", d, n, save)
	}
}