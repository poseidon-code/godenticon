package godenticon_test

type testIdenticonConfiguration struct {
    Size        []int
    Square      []bool
    Border      []bool
    Vertical    []bool
    Invert      []bool
    Symmetric   []bool
}
var tIC = testIdenticonConfiguration{
    Size:       []int{4,5,6,7,8},
    Square:     []bool{true, false},
    Border:     []bool{true, false},
    Vertical:   []bool{true, false},
    Invert:     []bool{true, false},
    Symmetric:  []bool{true, false},
}


type testImageConfiguration struct {
    Size        []string
    Portrait    []bool
    FG          []string
    BG          []string
}
var tImC = testImageConfiguration{
    Size:       []string{"S", "M", "L", "X"},
    Portrait:   []bool{true, false},
    FG:         []string{"03fcba"},
    BG:         []string{"013225"},
}
