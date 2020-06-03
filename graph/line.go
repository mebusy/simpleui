package graph

import (
    "math"
    // "image"
    "image/color"
    "image/draw"
)


func drawLineV( dst draw.Image, y1,  y2,  x1 float64, color color.Color ) {
    var x       int
    var y       int
    var nlen     int
    var cnt     int

    x = int(x1)
    y = int(math.Min( y1, y2 ))
    nlen = int(math.Abs( y1-y2 )) + 1

    cnt = 0
    for  cnt < nlen  {
        dst.Set( x,y+cnt,color )
        cnt = cnt + 1
    }
}

func drawLineH( dst draw.Image, x1,  x2,  y1 float64, color color.Color ) {
    var x       int
    var y       int
    var nlen     int
    var cnt     int

    y = int(y1)
    x = int(math.Min( x1, x2 ))
    nlen = int(math.Abs( x1-x2 )) + 1

    cnt = 0
    for  cnt < nlen  {
        dst.Set( x+cnt,y,color )
        cnt = cnt + 1
    }
}

func DrawLine( dst draw.Image, x1,  y1, x2, y2 float64 , color color.Color) {
    var x        int
    var y        int
    var dx       int
    var dy       int
    var dy_pos   int
    var a        int
    var b        int
    var diff     int


    if int(x1) == int(x2) {
        drawLineV( dst, y1, y2 , x1, color)
        return
    }
    if int(y1) == int(y2) {
        drawLineH( dst,x1, x2, y1, color)
        return ;
    }

    // fixed x,y on left-side
    if x1 < x2 {
        x = int(x1)
        y = int(y1)
        dx = int(x2 - x1)
        dy = int(y2 - y1)
        dy_pos = int(math.Abs(float64(dy)))
    } else {
        x = int(x2)
        y = int(y2)
        dx = int(x1 - x2)
        dy = int(y1 - y2)
        dy_pos = int(math.Abs(float64(dy)))
    }

    a = 0
    b = 0

    for {
        if a > dx {
            return
        }
        if b > dy_pos {
            return
        }

        if dy < 0 {
            dst.Set( x+a, y-b , color);
        } else {
            dst.Set( x+a, y+b, color )
        }

        if diff < 0 {
            a = a + 1
            diff = diff + dy_pos
        } else {
            b = b + 1
            diff = diff - dx
        }

    }

}

