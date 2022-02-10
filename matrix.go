package godenticon

// decides on putting 0/1 in every cell of matrix
func fill_cell(a, x, y int) int {
    if a % 2 != 0 {return x} else {return y}
}

// generates a symmetrically filled matrix
// i.e.: vertically half of the matrix is mirror opposite of other half
func (i *Identicon) generate_symmetric_matrix(bit_0, bit_1 int) {
    w, h := (i.Width*i.IdenticonOptions.Size), (i.Height*i.IdenticonOptions.Size)
    m := make([][]int, h)
    for r:=0; r<h; r++ {m[r] = make([]int, w)}

    k, b := 1, 0
    for r:=0; r<h; r++ {
        for c:=0; c<w; c++ {
            if c>=w/2+1 {k++; continue}
            b = fill_cell(int(i.Hash[k-1]), bit_1, bit_0)
            m[r][c], m[r][w-c-1] = b, b
            k++
        }
    }

    i.Matrix = m
}

// generates an asymmetrically filled matrix
func (i *Identicon) generate_asymmetric_matrix(bit_0, bit_1 int) {
    w, h := (i.Width*i.IdenticonOptions.Size), (i.Height*i.IdenticonOptions.Size)
    m := make([][]int, h)
    for r:=0; r<h; r++ {m[r] = make([]int, w)}

    k, b := 1, 0
    for r:=0; r<h; r++ {
        for c:=0; c<w; c++ {
            b = fill_cell(int(i.Hash[k-1]), bit_1, bit_0)
            m[r][c] = b
            k++
        }
    }

    i.Matrix = m
}


func (i *Identicon) GenerateMatrix() {
    // bits for matrix cell filling
    bit_0, bit_1 := 0, 1

    if i.IdenticonOptions.Vertical {
        if i.IdenticonOptions.Invert {
            // bits are switched, hence matrix cell filling is also switched (inversion)
            bit_0, bit_1 = 1, 0

            if i.IdenticonOptions.Symmetric {
                // creates a matrix of 0s (empty cells) & 1s (filled cells), where:
                // - the aspect ratio is switched
                // - cell filling is switched
                // - cell filling is symmetric
                i.generate_symmetric_matrix(bit_0, bit_1)
            } else {
                // creates a matrix of 0s (empty cells) & 1s (filled cells), where:
                // - the aspect ratio is switched
                // - cell filling is switched
                i.generate_asymmetric_matrix(bit_0, bit_1)
            }
        } else {
            if i.IdenticonOptions.Symmetric {
                // creates a matrix of 0s (empty cells) & 1s (filled cells), where:
                // - the aspect ratio is switched
                // - cell filling is symmetric
                i.generate_symmetric_matrix(bit_0, bit_1)
            } else {
                // creates a matrix of 0s (empty cells) & 1s (filled cells), where:
                // - the aspect ratio is switched
                i.generate_asymmetric_matrix(bit_0, bit_1)
            }
        }
    } else {
        if i.IdenticonOptions.Invert {
            // bits are switched, hence matrix cell filling is also switched (inversion)
            bit_0, bit_1 = 1, 0

            if i.IdenticonOptions.Symmetric {
                // creates a matrix of 0s (empty cells) & 1s (filled cells), where:
                // - cell filling is switched
                // - cell filling is symmetric
                i.generate_symmetric_matrix(bit_0, bit_1)
            } else {
                // creates a matrix of 0s (empty cells) & 1s (filled cells), where:
                // - cell filling is switched
                i.generate_asymmetric_matrix(bit_0, bit_1)
            }
        } else {
            if i.IdenticonOptions.Symmetric {
                // creates a matrix of 0s (empty cells) & 1s (filled cells), where:
                // - cell filling is symmetric
                i.generate_symmetric_matrix(bit_0, bit_1)
            } else {
                // creates a matrix of 0s (empty cells) & 1s (filled cells)
                i.generate_asymmetric_matrix(bit_0, bit_1)
            }
        }
    }
}
