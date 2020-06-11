package graph

import (
    "image"
    "image/color"
    "image/draw"
)

var (
    COLOR_RED = color.RGBA{255,0,0,255}
    COLOR_LIME = color.RGBA{0,255,0,255}
    COLOR_BLUE = color.RGBA{0,0,255,255}
    COLOR_YELLOW = color.RGBA{255,255,0,255}
    COLOR_CYAN = color.RGBA{0,255,255,255}
    COLOR_MAGENTA = color.RGBA{255,0,255,255}

    COLOR_SILVER = color.RGBA{192,192,192,255}
    COLOR_GRAY = color.RGBA{128,128,128,255}
    COLOR_MAROON = color.RGBA{128,0,0,255}
    COLOR_OLIVE = color.RGBA{128,128,0,255}

    COLOR_GREEN = color.RGBA{0,128,0,255}
    COLOR_PURPLE = color.RGBA{128,0,128,255}
    COLOR_TEAL = color.RGBA{0,128,128,255}
    COLOR_NAVY = color.RGBA{0,0,128,255}

    COLOR_DARKBLUE = color.RGBA{0,0,139,255}
    COLOR_DARKGREEN = color.RGBA{1,50,32,255}
    COLOR_DARKGRAY = color.RGBA{169,169,169,255}
    COLOR_DIMGRAY = color.RGBA{105,105,105,255}
)

func CopyStride( dst_buf []uint8, dst_stride int ,  src_buf []uint8, src_stride int , nBytePerLine int , nLines int ) {
    dst := 0
    src := 0
    for h:=0; h< nLines; h++ {
        copy( dst_buf[dst:dst+nBytePerLine],  src_buf[src:src+nBytePerLine ] )
        dst += dst_stride
        src += src_stride
    }
}

func FillRect( dst draw.Image , r image.Rectangle , c color.Color ) {
    draw.Draw( dst, r, &image.Uniform{c}, image.ZP, draw.Src  )
}


