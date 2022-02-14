package godenticon

import (
	"encoding/hex"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type image_dimension struct { w, h int }
var ds = image_dimension{1920, 1080}
var dm = image_dimension{2560, 1440}
var dl = image_dimension{3840, 2160}
var dx = image_dimension{7680, 4320}

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
        fmt.Println("Invalid image size.") 
        fmt.Println("Image size (string) value should be any one of S, M, L & X.")
        fmt.Println("i.e.: Identicon.ImageOptions.Size='X'")
        os.Exit(1)
    }

    return -1, -1
}

func hex_to_rgb(h string) color.Color {
    if len(h)!=6 {
        fmt.Println("Color should be in HEX format of length 6 (range: '000000' to 'ffffff')")
        os.Exit(1)
    }

    rgb, err := hex.DecodeString(h)
    if err!=nil {
        fmt.Println("Invalid HEX color :", h)
        fmt.Println(err)
        os.Exit(1)
    }

    return color.RGBA{rgb[0], rgb[1], rgb[2], 255}
}


func (i *Identicon) SaveImage(path string) {
    mw, mh := len(i.Matrix[0]), len(i.Matrix)               // matrix width(#columns) & height(#rows)
    iw, ih := get_image_dimension(i.ImageOptions.Size)      // image width & height (in pixels)

    // TODO: handle portrait image

    // TODO: appropriate cell-size `b` calculations
    b := ih/6/mh        // identicon cell size (in pixels)
    bw := b/2           // border thickness (border width)
    
    img := image.NewRGBA(image.Rect(0,0,iw,ih))     // image canvas
    fg := hex_to_rgb(i.ImageOptions.FG)             // image foreground color
    bg := hex_to_rgb(i.ImageOptions.BG)             // image backgoround color

    // set background
    for x:=0; x<iw; x++ {for y:=0; y<ih; y++ {img.Set(x, y, bg)}}

    // centering identicon coordinates (offsets)
    ox := (iw/2) - (mw*b/2)
    oy := (ih/2) - (mh*b/2)


    // set border
    if i.IdenticonOptions.Border {
        xs, ys := ox-(3*bw), oy-(3*bw)
        xe, ye := mw*b+ox+(3*bw), mh*b+oy+(3*bw)
    
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
                for x:=(ox+(c*b)); x<(ox+(c*b))+b; x++ {
                    for y:=(oy+(r*b)); y<(oy+(r*b))+b; y++ {
                        img.Set(x, y, fg)
                    }
                }
            }
        }
    }

    
    save := handle_save_path(path, i.Text, ".png")
    f, _ := os.Create(save); defer f.Close()
    png.Encode(f, img)
}