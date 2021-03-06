package graph

import (
    "math"
    // "image"
    "image/color"
    "image/draw"
)


func drawLineV( dst draw.Image, y1,  y2,  x1 int, color color.Color ) {
    var x       int
    var y       int
    var nlen     int
    var cnt     int

    x = x1
    y = int(math.Min( float64(y1), float64(y2) ))
    nlen = int(math.Abs( float64(y1-y2) )) + 1

    cnt = 0
    for  cnt < nlen  {
        dst.Set( x,y+cnt,color )
        cnt = cnt + 1
    }
}

func drawLineH( dst draw.Image, x1,  x2,  y1 int, color color.Color ) {
    var x       int
    var y       int
    var nlen     int
    var cnt     int

    y = y1
    x = int(math.Min( float64(x1), float64(x2) ))
    nlen = int(math.Abs( float64(x1-x2) )) + 1

    cnt = 0
    for  cnt < nlen  {
        dst.Set( x+cnt,y,color )
        cnt = cnt + 1
    }
}

func DrawLine( dst draw.Image, x1,  y1, x2, y2 int , color color.Color) {
    var x        int
    var y        int
    var dx       int
    var dy       int
    var dx_pos   int
    var dy_pos   int
    var a        int
    var b        int
    var diff     int

    /*
    x1 := int(_x1)
    y1 := int(_y1)
    x2 := int(_x2)
    y2 := int(_y2)
    */

    if x1 == x2 {
        drawLineV( dst, y1, y2 , x1, color)
        return
    }
    if y1 == y2 {
        drawLineH( dst,x1, x2, y1, color)
        return
    }

    x = x1
    y = y1
    dx = x2 - x1
    dy = y2 - y1
    dx_pos = int(math.Abs(float64(dx)))
    dy_pos = int(math.Abs(float64(dy)))

    a = 0
    b = 0

    for {
        if a > dx_pos {
            return
        }
        if b > dy_pos {
            return
        }

        if dx < 0 && dy < 0 {
            dst.Set( x-a, y-b , color)
        } else if dx > 0 && dy < 0 {
            dst.Set( x+a, y-b , color)
        } else if dx < 0 && dy > 0 {
            dst.Set( x-a, y+b , color)
        } else {
            dst.Set( x+a, y+b, color )
        }

        if diff < 0 {
            a = a + 1
            diff = diff + dy_pos
        } else {
            b = b + 1
            diff = diff - dx_pos
        }

    }

}

