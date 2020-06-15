package graph

import (
    // "math"
    // "image"
    "image/color"
    "image/draw"
    "sort"
    // "log"
)

type Point2D struct {
    x,y int
}

type Triangle2D struct {
    verts []Point2D
}

func (self *Triangle2D) SortByY() {
    sort.SliceStable(self.verts, func(i, j int) bool {
        return self.verts[i].y < self.verts[j].y
    })
}
func (self *Triangle2D) GetVert( i int ) Point2D {
    return self.verts[i]
}

func (self *Triangle2D) SetVert( i int, x,y int ) {
    self.verts[i].x = x
    self.verts[i].y = y
}

func NewTriangle(x1,y1,x2,y2,x3,y3 int ) Triangle2D  {
    return Triangle2D{ []Point2D { {x1,y1},{x2,y2},{x3,y3}  } }
}

func fillBottomFlatTriangle( dst draw.Image,
    x1,y1, x2,y2, x3,y3 int, color color.Color ) {

    invslope1 := float64(x2 - x1) / float64(y2 - y1)
    invslope2 := float64(x3 - x1) / float64(y3 - y1)

    curx1 := float64(x1)
    curx2 := float64(x1)

    for scanlineY := y1; scanlineY <= y2; scanlineY++ {
        drawLineH(dst, int(curx1), int(curx2), scanlineY, color)
        curx1 += invslope1
        curx2 += invslope2
    }
}

func fillTopFlatTriangle( dst draw.Image,
    x1,y1, x2,y2, x3,y3 int, color color.Color ) {

    invslope1 := float64(x3 - x1) / float64(y3 - y1)
    invslope2 := float64(x3 - x2) / float64(y3 - y2)

    curx1 := float64(x3)
    curx2 := float64(x3)

    for scanlineY := y3; scanlineY > y1; scanlineY-- {
        drawLineH(dst, int(curx1), int(curx2), scanlineY, color)
        curx1 -= invslope1
        curx2 -= invslope2
    }
}

func FillTriangle( dst draw.Image, triangle Triangle2D,  color color.Color ) {
   /* at first sort the three vertices by y-coordinate ascending so v1 is the topmost vertice */
   triangle.SortByY()

   v1 := triangle.GetVert(0)
   v2 := triangle.GetVert(1)
   v3 := triangle.GetVert(2)

   /* here we know that y1 <= y2 <= y3 */
   /* check for trivial case of bottom-flat triangle */
   if v2.y == v3.y {
       fillBottomFlatTriangle(dst, v1.x, v1.y, v2.x, v2.y, v3.x,v3.y, color)
   } else if v1.y == v2.y {
       /* check for trivial case of top-flat triangle */
       fillTopFlatTriangle(dst, v1.x, v1.y, v2.x, v2.y, v3.x,v3.y, color)
   } else {
       /* general case - split the triangle in a topflat and bottom-flat one */
       v4 := Point2D{ int(float64(v1.x) + (float64(v2.y - v1.y) / float64(v3.y - v1.y)) * float64(v3.x - v1.x)), v2.y }
       fillBottomFlatTriangle(dst, v1.x,v1.y, v2.x,v2.y, v4.x,v4.y , color );
       fillTopFlatTriangle(dst, v2.x,v2.y, v4.x,v4.y, v3.x,v3.y , color );
   }
}


func DrawTriangle( dst draw.Image, triangle Triangle2D,  color color.Color ) {
    v1 := triangle.GetVert(0)
    v2 := triangle.GetVert(1)
    v3 := triangle.GetVert(2)
    DrawLine( dst, v1.x,v1.y, v2.x,v2.y, color  )
    DrawLine( dst, v2.x,v2.y, v3.x,v3.y, color  )
    DrawLine( dst, v3.x,v3.y, v1.x,v1.y, color  )
}
