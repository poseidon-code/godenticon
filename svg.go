package godenticon

import (
	"fmt"
	"os"
)

// Creates and saves an identicon as SVG. Requires a `path` variable
// to be passed, as a SVG saving directory or name. Size of the SVG
// identicon depends on the Cell Size of every block in the identicon,
// which is currently hardcoded to 100 units.
func (i *Identicon) SaveSVG(path string) {
    bs  := 10       // identicon block size
    bw  := 0        // border(stroke) width(thickness)
    ip  := 0        // identicon padding (when using border)

    mw, mh := len(i.Matrix[0]), len(i.Matrix)       // matrix width & height
    sw, sh := mw*bs, mh*bs                          // SVG width & height

    if i.IdenticonOptions.Border {
        bw = bs/3
        ip = bs
        sw += bw*2+ip
        sh += bw*2+ip
    }

    svg := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 %d %d" width="%d" height="%d" shape-rendering="crispEdges">`, sw, sh, sw, sh)
    if i.IdenticonOptions.Border {
        svg += fmt.Sprintf(`<rect width="%d" height="%d" stroke="black" stroke-width="%d" fill-opacity="0" />`, sw, sh, 2*bw)
    }

    for r:=0; r<mh; r++ {
        svg += "<g>"
        for c:=0; c<mw; c++ {
            if i.Matrix[r][c] == 1 {
                svg += fmt.Sprintf(`<rect width="%d" height="%d" x="%d" y="%d" />`, bs, bs, (c*bs)+bw+(ip/2), (r*bs)+bw+(ip/2))
            }
        }
        svg += "</g>"
    }
    svg += "</svg>"
    
    save := handleSavePath(path, "svg")
    f, _ := os.Create(save); defer f.Close()
    f.WriteString(svg)
}