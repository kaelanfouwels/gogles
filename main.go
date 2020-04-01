package main // import "github.com/kaelanfouwels/gogles"

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/kaelanfouwels/gogles/fontman"
	"github.com/kaelanfouwels/gogles/textman"

	"github.com/go-gl/glfw/v3.3/glfw"
	gl "github.com/kaelanfouwels/gogles/glow/gl"
	"github.com/kaelanfouwels/gogles/renderman"
	//gl "github.com/kaelanfouwels/gogles/glow/gles"
)

const width, height = 800, 600

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	err := start()
	if err != nil {
		log.Printf("%v", err)
		os.Exit(1)
	}
}

func start() error {
	if err := glfw.Init(); err != nil {
		return fmt.Errorf("failed to initialize glfw: %w", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	window, err := glfw.CreateWindow(width, height, "gogles", nil, nil)
	if err != nil {
		return err
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		return err
	}

	textman, err := textman.NewTextman("./assets")
	if err != nil {
		return err
	}
	defer textman.Destroy()

	fontman, err := fontman.NewFontman(textman)
	if err != nil {
		return err
	}

	renderman, err := renderman.NewRenderman(width, height, textman, fontman)
	if err != nil {
		return err
	}
	defer renderman.Destroy()

	for !window.ShouldClose() {
		err := renderman.Draw()
		if err != nil {
			return fmt.Errorf("Draw cycle failed: %w", err)
		}
		window.SwapBuffers()
		glfw.PollEvents()
	}

	return nil
}
