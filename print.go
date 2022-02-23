package godenticon

import "fmt"

// Printing the Identicon.Matrix to the terminal.
// Identicon.IdenticonOptions.Border (bool) is used to determine whether
// to print the identicon with/without border.
func (i *Identicon) Print() {
    w, h := len(i.Matrix[0]), len(i.Matrix)
    m := i.Matrix
    
    if i.IdenticonOptions.Border {
        w, h = w+4, h+2

        for r:=0; r<h; r++ {
            for c:=0; c<w; c++ {
                if r==0 {
                    if c==0 {
                        fmt.Print("⎡")
                    } else if c==w-1 {
                        fmt.Print("⎤")
                    } else {
                        fmt.Print("⎺⎺")
                    }
                } else if r==h-1 {
                    if c==0 {
                        fmt.Print("⎣")
                    } else if c==w-1 {
                        fmt.Print("⎦")
                    } else {
                        fmt.Print("__")
                    }
                } else {
                    if c==0 {
                        fmt.Print("⎢")
                    } else if c==w-1 {
                        fmt.Print("⎥")
                    } else if c==1 || c==w-2 {
                        fmt.Print("  ")
                    } else {
                        if m[r-1][c-2] == 0 {
                            fmt.Print("  ")
                        } else {
                            fmt.Print("██")
                        }
                    }
                }
            }
            fmt.Println()
        }
    } else {
        for r:=0; r<h; r++ {
            for c:=0; c<w; c++ {
                if m[r][c] == 0 {
                    fmt.Print("  ")
                } else {
                    fmt.Print("██")
                }
            }
            fmt.Println()
        }
    }
}