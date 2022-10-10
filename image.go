package godenticon

import (
	"encoding/hex"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

type image_dimension struct { w, h int }
var ds = image_dimension{1920, 1080}
var dm = image_dimension{2560, 1440}
var dl = image_dimension{3840, 2160}
var dx = image_dimension{7680, 4320}

// get appropriate image dimensions from size string
func get_image_dimension(s string) (w, h int) {
    if s=="S" || s=="M" || s=="L" || s=="X" {
        switch s {
        case "S":
            return ds.w, ds.h
        case "M":
            return dm.w, dm.h
        case "L":
            return dl.w, dl.h
        case "X":
            return dx.w, dx.h
        }
    } else {
        log.Fatalln(
            "Invalid image size:", s, 
            "\nImage size (string) value should be any one of S, M, L & X.",
            "\ni.e.: Identicon.ImageOptions.Size='X'",
        )
    }

    return -1, -1
}

// HEX color to RGB color conversion
func hex_to_rgb(h string) color.Color {
    rgb, _ := hex.DecodeString(h)
    return color.RGBA{rgb[0], rgb[1], rgb[2], 255}
}

// appropriate block-size `bs` calculation
func getBlockSize(imageFactor, matrixFactor int, s, p, v bool) (bs int) {
    factor := imageFactor/matrixFactor

    if s {
        if p { bs = factor/10
        } else { bs = factor/5 }
    } else {
        if p {
            if v { bs = factor/5
            } else { bs = factor/10 }
        } else {
            if v { bs = factor/3
            } else { bs = factor/6 }
        }
    }

    return bs
}

// appropriate border-width `bw` calculation
func get_border_width(mw, mh, bs int) (bw int) {
    if mh==4 || mw==4 {
        bw = bs/4
    } else if mh<=6 || mw<=6 {
        bw = bs/3
    } else {
        bw = bs/2
    }

    return bw
}


// Creates and saves an identicon as PNG image. Requires a `path` variable
// to be passed, as an image saving directory or name. Size of the image
// identicon depends on the Identicon.ImageOptions.Size property which is minimum
// of either width & height of the image, divided by some value.
func (i *Identicon) SaveImage(path string) {
    mw, mh := len(i.Matrix[0]), len(i.Matrix)               // matrix width(#columns) & height(#rows)
    iw, ih := get_image_dimension(i.ImageOptions.Size)      // image width & height (in pixels)

    // handle portrait image
    if i.ImageOptions.Portrait {
        iw, ih = ih, iw
    }

    // identicon cell size (in pixels)
    bs := getBlockSize(ih, mh, i.IdenticonOptions.Square, i.ImageOptions.Portrait, i.IdenticonOptions.Vertical)
    // border thickness (border width)
    bw := get_border_width(mw, mh, bs)
    
    img := image.NewRGBA(image.Rect(0,0,iw,ih))     // image canvas
    fg := hex_to_rgb(i.ImageOptions.FG)             // image foreground color
    bg := hex_to_rgb(i.ImageOptions.BG)             // image backgoround color

    // set background
    for x:=0; x<iw; x++ {for y:=0; y<ih; y++ {img.Set(x, y, bg)}}

    // centering identicon coordinates (offsets)
    ox := (iw/2) - (mw*bs/2)
    oy := (ih/2) - (mh*bs/2)


    // set border
    if i.IdenticonOptions.Border {
        xs, ys := ox-(3*bw),        oy-(3*bw)
        xe, ye := ox+(mw*bs)+(3*bw), oy+(mh*bs)+(3*bw)
    
        br := [][]int{
            {xs,    xe,     ys,     ys+bw},
            {xs,    xs+bw,  ys,     ye},
            {xs,    xe,     ye-bw,  ye},
            {xe-bw, xe,     ys,     ye},
        }

        for _, v := range br {
            for x:=v[0]; x<v[1]; x++ {
                for y:=v[2]; y<v[3]; y++ {
                    img.Set(x, y, fg)
                }
            }
        }
    }


    // set identicon blocks
    for r:=0; r<mh; r++ {
        for c:=0; c<mw; c++ {
            if i.Matrix[r][c] == 1 {
                for x:=(ox+(c*bs)); x<(ox+(c*bs))+bs; x++ {
                    for y:=(oy+(r*bs)); y<(oy+(r*bs))+bs; y++ {
                        img.Set(x, y, fg)
                    }
                }
            }
        }
    }


    save := handleSavePath(path, "png")
    f, _ := os.Create(save); defer f.Close()
    png.Encode(f, img)
}
