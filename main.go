package main

import "C"
import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/faiface/mainthread"
	"github.com/go-gl/gl/v3.3-core/gl"
	// "strings"
	"unsafe"
	"fmt"
	"time"
	"math/rand"
	"github.com/freegamerskids/totallynotanstartingoutprojectforgolangpleasedontlookatmycodeok/settings"
)

settings.RED = *0;
settings.GREEN = *0;
settings.BLUE = *0;
settings.ALPHA = *1;

func run() {


	mainthread.Call(func(){
		glfw.Init()
		glfw.WindowHint(glfw.ContextVersionMajor, 3)
		glfw.WindowHint(glfw.ContextVersionMinor, 3)
		glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
		glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
		glfw.WindowHint(glfw.Resizable, glfw.True)
		glfw.WindowHint(glfw.Samples, 0)

		monitor := glfw.GetPrimaryMonitor()
		mWidth, mHeight := monitor.GetVideoMode().Width, monitor.GetVideoMode().Height

		fmt.Println("Monitor pixels: ", string(int(mWidth)), " x ", string(int(mHeight)))

		glfw.WindowHint(glfw.RedBits, monitor.GetVideoMode().RedBits)
		glfw.WindowHint(glfw.GreenBits, monitor.GetVideoMode().GreenBits)
		glfw.WindowHint(glfw.BlueBits, monitor.GetVideoMode().BlueBits)
		glfw.WindowHint(glfw.RefreshRate, monitor.GetVideoMode().RefreshRate)
		//glfw.WindowHint(glfw.Decorated, glfw.False)
		win, err := glfw.CreateWindow(int(mWidth), int(mHeight), "test", monitor, nil);

		if err != nil {
			panic(err)
		}

		win.SetTitle("testing window")

		win.MakeContextCurrent()
		win.SetShouldClose(false)
		win.Focus()
		fmt.Println("clipboard: ", win.GetClipboardString())

		fmt.Println("GLFW initialized!")

		gl.Init()

		C.GoString((*C.char)(unsafe.Pointer(gl.GetString(gl.RENDERER))))

		glVendor := C.GoString((*C.char)(unsafe.Pointer(gl.GetString(gl.VENDOR))))
		glRenderer := C.GoString((*C.char)(unsafe.Pointer(gl.GetString(gl.RENDERER))))
		glVersion := C.GoString((*C.char)(unsafe.Pointer(gl.GetString(gl.VERSION))))
		glslVersion := C.GoString((*C.char)(unsafe.Pointer(gl.GetString(gl.SHADING_LANGUAGE_VERSION))))

		extensions := ""

		numExtensions := 0
		gl.GetIntegerv(gl.NUM_EXTENSIONS, &numExtensions)

		for i := int32(0); i < numExtensions; i++ {
			extensions += C.GoString((*C.char)(unsafe.Pointer(gl.GetStringi(gl.EXTENSIONS, uint32(i)))))
			extensions += " "
		}

		fmt.Println("GL Vendor:    ", glVendor)
		fmt.Println("GL Renderer:  ", glRenderer)
		fmt.Println("GL Version:   ", glVersion)
		fmt.Println("GLSL Version: ", glslVersion)
		fmt.Println("GL Extensions:", extensions)
		fmt.Println("OpenGL initialized!")

		for !win.ShouldClose() {
			pressedP := false
			mainthread.Run(mainthread.Call(func(){
				setCol(Uint8(rand), Uint8(rand), Uint8(rand), rand.Float32)
				pushFrame()

				if win.GetKey(glfw.KeyEscape) == glfw.Press {
					win.SetShouldClose(true)
				}

				if win.GetKey(glfw.KeyP) == glfw.Press {

					if pressedP {
						fmt.Println("pressed")
					}

					pressedP = true
				}

				if win.GetKey(glfw.KeyP) == glfw.Release {
					pressedP = false
				}

				win.SwapBuffers()

			}))
		}
	})
}

func pushFrame() {
	glfw.PollEvents()

	gl.Enable(gl.SCISSOR_TEST)
	gl.Disable(gl.DITHER)

	gl.ClearColor(settings.R, settings.G, settings.B, settings.A)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func setCol(uint8 r, uint8 g, uint8 b, uint8 a) {
	settings.R = r
	settings.G = g 
	settings.B = b
	settings.A = a
}

func sleep(int secs){
	tosleep := secs * time.Second

	time.Sleep(tosleep)
}

func randIntRange(r *rand.Rand, min, max int) int {
	if min == max {
		return min
	}
	return r.Intn((max+1)-min) + min
}

func (f *Faker) Uint8() uint8 { return uint8Func(f.Rand) }

func uint8Func(r *rand.Rand) uint8 { return uint8(randIntRange(r, 0, math.MaxUint8)) }

func main() {
	mainthread.Run(run)
}