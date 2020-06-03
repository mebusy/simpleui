package simpleui

import (
	"crypto/md5"
	"path"
	//"encoding/binary"
	"fmt"
	"image"

	// "image/color"
	"image/draw"
	// "image/gif"
	"image/png"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"os/user"

	//"path"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	// "nes"
)

var homeDir string

func init() {
    u, err := user.Current()
    if err != nil {
        log.Fatalln(err)
    }
    homeDir = u.HomeDir
}

func HomeDir() string {
    return homeDir
}


func ReadKey(window *glfw.Window, key glfw.Key) bool {
    return window.GetKey(key) == glfw.Press
}


func readJoystick(joy glfw.Joystick, turbo bool) [8]bool {
    var result [8]bool
    /*
        if !glfw.JoystickPresent(joy) {
            return result
        }
        joyname := glfw.GetJoystickName(joy)
        axes := glfw.GetJoystickAxes(joy)
        buttons := glfw.GetJoystickButtons(joy)
        if joyname == "PLAYSTATION(R)3 Controller" {
            result[nes.ButtonA] = buttons[14] == 1 || (turbo && buttons[2] == 1)
            result[nes.ButtonB] = buttons[13] == 1 || (turbo && buttons[3] == 1)
            result[nes.ButtonSelect] = buttons[0] == 1
            result[nes.ButtonStart] = buttons[3] == 1
            result[nes.ButtonUp] = buttons[4] == 1 || axes[1] < -0.5
            result[nes.ButtonDown] = buttons[6] == 1 || axes[1] > 0.5
            result[nes.ButtonLeft] = buttons[7] == 1 || axes[0] < -0.5
            result[nes.ButtonRight] = buttons[5] == 1 || axes[0] > 0.5
            return result
        }
        if len(buttons) < 8 {
            return result
        }
        result[nes.ButtonA] = buttons[0] == 1 || (turbo && buttons[2] == 1)
        result[nes.ButtonB] = buttons[1] == 1 || (turbo && buttons[3] == 1)
        result[nes.ButtonSelect] = buttons[6] == 1
        result[nes.ButtonStart] = buttons[7] == 1
        result[nes.ButtonUp] = axes[1] < -0.5
        result[nes.ButtonDown] = axes[1] > 0.5
        result[nes.ButtonLeft] = axes[0] < -0.5
        result[nes.ButtonRight] = axes[0] > 0.5
        //*/
    return result
}

func joystickReset(joy glfw.Joystick) bool {
    if !glfw.JoystickPresent(joy) {
        return false
    }
    buttons := glfw.GetJoystickButtons(joy)
    if len(buttons) < 6 {
        return false
    }
    return buttons[4] == 1 && buttons[5] == 1
}


func HashFile(path string) (string, error) {
    data, err := ioutil.ReadFile(path)
    if err != nil {
        return "", err
    }
    return fmt.Sprintf("%x", md5.Sum(data)), nil
}

func createGLTexture() uint32 {
    var texture uint32
    gl.GenTextures(1, &texture)
    gl.BindTexture(gl.TEXTURE_2D, texture)
    gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
    gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
    gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
    gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
    gl.BindTexture(gl.TEXTURE_2D, 0)
    return texture
}

func setTexture(im *image.RGBA) {
    size := im.Rect.Size()
    gl.TexImage2D(
        gl.TEXTURE_2D, 0, gl.RGBA, int32(size.X), int32(size.Y),
        0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(im.Pix))
}

func copyImage(src image.Image) *image.RGBA {
    dst := image.NewRGBA(src.Bounds())
    draw.Draw(dst, dst.Rect, src, image.ZP, draw.Src)
    return dst
}

func loadPNG(path string) (image.Image, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    return png.Decode(file)
}

func savePNG(path string, im image.Image) error {
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()
    return png.Encode(file, im)
}

func saveJPG(path string, im image.Image) error {
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()
    return jpeg.Encode(file, im, &jpeg.Options{ Quality:100 } )
}


func saveGIF(path string, frames []image.Image) error {
    /*
        var palette []color.Color
        for _, c := range nes.Palette {
            palette = append(palette, c)
        }
        g := gif.GIF{}
        for i, src := range frames {
            if i%3 != 0 {
                continue
            }
            dst := image.NewPaletted(src.Bounds(), palette)
            draw.Draw(dst, dst.Rect, src, image.ZP, draw.Src)
            g.Image = append(g.Image, dst)
            g.Delay = append(g.Delay, 5)
        }
        file, err := os.Create(path)
        if err != nil {
            return err
        }
        defer file.Close()
        return gif.EncodeAll(file, &g)
        //*/
    return nil
}

func Screenshot(directory string, im image.Image) {
    os.MkdirAll(directory, os.ModePerm)
    for i := 0; i < 256; i++ {
        path :=  path.Join(directory,  fmt.Sprintf("%03d.jpg", i))
        if _, err := os.Stat(path); os.IsNotExist(err) {
            saveJPG(path, im)
            log.Println("save snaphost to" , path )
            return
        }
    }
}

func animation(frames []image.Image) {
    for i := 0; i < 1000; i++ {
        path := fmt.Sprintf("%03d.gif", i)
        if _, err := os.Stat(path); os.IsNotExist(err) {
            saveGIF(path, frames)
            return
        }
    }
}

