package main

import (
    "github.com/mebusy/simpleui"
    "github.com/go-gl/glfw/v3.1/glfw"
    "image"
    "image/color"
    "github.com/mebusy/simpleui/graph"
    // "image/draw"
    "math/rand"
    "log"
    "time"
)

type MyView struct {
    screenImage *image.RGBA
}

func NewView( w,h int) *MyView {
    view := &MyView{}
    view.screenImage = image.NewRGBA(image.Rect(0, 0, w, h))
    return view
}

var test_pts [6]int
func (self *MyView) Enter() {
    rand.Seed( time.Now().Unix() )
    for i:=0; i< len(test_pts); i++ {
        test_pts[i] = rand.Intn(200)
    }
    log.Println( "draw triangle:" , test_pts )
}
func (self *MyView) Exit() {}
func (self *MyView) Update(t, dt float64) {

    graph.FillRect( self.screenImage, self.screenImage.Bounds() ,
                graph.COLOR_CYAN )

    self.screenImage.Set( 10,10, color.White )
    self.screenImage.Set( 11,11, color.White )

    graph.DrawLine( self.screenImage,  -10, -20 , 100,160,  graph.COLOR_RED )
    graph.DrawLine( self.screenImage,  100, -20 , 0,100,  graph.COLOR_GREEN )

    graph.DrawLine( self.screenImage,   105,160, -5, -20 , graph.COLOR_BLUE )
    graph.DrawLine( self.screenImage,   5,100, 105, -20 , graph.COLOR_YELLOW )

    graph.DrawLine( self.screenImage,14.358431036073966,32.12406230496519,24.339200799863534,31.504195446499647,  color.Black )

    graph.DrawTriangle(self.screenImage, graph.NewTriangle(test_pts[0],test_pts[1],test_pts[2],test_pts[3],test_pts[4],test_pts[5])  , graph.COLOR_GREEN)

    graph.DrawRectangle( self.screenImage, 10, 200,  10+50, 200-80, graph.COLOR_RED )

    graph.FillCircle( self.screenImage, 290, 50, 30, graph.COLOR_YELLOW )
    graph.DrawCircle( self.screenImage, 300, 50, 30, graph.COLOR_BLUE )

}

func (self *MyView) OnKey(key glfw.Key) {}
func (self *MyView) TextureBuff() []uint8 {
    return self.screenImage.Pix
}
func (self *MyView) Title() string {
    return "my game"
}


func main() {
    w,h,scale := 320,240,2
    view := NewView(w,h)
    simpleui.SetWindow( w,h, scale  )
    simpleui.Run( view )
}
