package graph

import (
    "image/draw"
    "image/color"
)

/** Draws a filled rectangle whose top left corner is (x1, y1)
 * and bottom right corner is (x2,y2), using the current color. 
 */
func FillRectangle( dst draw.Image, x1,  y1,  x2, y2 int, color color.Color) {
    var  y, y_end int

    if y1 < y2 {
        y = y1
        y_end = y2
    } else {
        y = y2
        y_end = y1
    }

    for  y < ( y_end + 1 ) {
        drawLineH( dst,  x1, x2, y, color )
        y = y + 1
    }
}

func DrawRectangle( dst draw.Image, x1,  y1,  x2, y2 int, color color.Color) {
    drawLineH( dst, x1,x2,y1, color )
    drawLineH( dst, x1,x2,y2, color )
    drawLineV( dst, y1,y2,x1, color )
    drawLineV( dst, y1,y2,x2, color )
}
