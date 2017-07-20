package main

import (
	"runtime"

	"log"

	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	width  = 512
	height = 512
)

func initEnv() {
	runtime.LockOSThread()
	window := initWindow()
	defer glfw.Terminate()

	for !window.ShouldClose() {

	}
}

func initWindow() *glfw.Window {
	err := glfw.Init()
	if err != nil {
		log.Panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "This is a demo.", nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	window.MakeContextCurrent()
	return window
}

func main() {
	initEnv()
	initWindow()
}
