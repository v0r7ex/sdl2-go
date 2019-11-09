package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

var running = true

func main() {
	// Initialize SDL
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Printf("Could not initialize SDL2: %s\n", sdl.GetError())
		return
	}
	defer sdl.Quit()

	// Create the window
	window, err := sdl.CreateWindow("SDL2", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 1024, 576, sdl.WINDOW_RESIZABLE)
	if err != nil {
		fmt.Printf("Could not create window: %s\n", sdl.GetError())
	}
	defer window.Destroy()

	// Run game loop
	for running {
		getEvents()
	}
}

func getEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		return
	}
}
