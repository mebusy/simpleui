package graph

import (
    "image/draw"
    "image/color"
    "math"
)

func drawCircle( dst draw.Image,  x, y, r int, color color.Color , bFill bool ) {
    var dx,dy int

    // left right edge of each drawing scanline
    xl := x
    xr := x
    dy = -r
    for  dy < ( r+1 ) {
        dx = int(math.Sqrt( float64(r*r) - float64(dy*dy) ))
        if bFill {
            drawLineH( dst, x-dx, x+dx, y + dy , color )
        } else {
            drawLineH( dst, x-dx , xl , y + dy , color )
            drawLineH( dst, x+dx , xr , y + dy , color )
            xl = x-dx
            xr = x+dx
        }
        dy = dy + 1
    }
}


func DrawCircle( dst draw.Image,  x, y, r int, color color.Color ) {
    drawCircle( dst, x, y, r, color, false )
}

func FillCircle( dst draw.Image,  x, y, r int, color color.Color ) {
    drawCircle( dst, x, y, r, color, true )
}
