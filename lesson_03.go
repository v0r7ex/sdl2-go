package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Game struct {
	running bool
	score int
	colored bool
}

func main() {
	game := Game {running: true}

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

	// Create renderer
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Printf("Could not create renderer: %s\n", sdl.GetError())
	}
	defer renderer.Destroy()

	// Run game loop
	for game.running {
		getEvents(&game)
		display(&game, renderer)
	}
}

func getEvents(game *Game) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			game.running = false
			break
		case *sdl.KeyboardEvent:
			if e.Type == sdl.KEYDOWN {
				switch e.Keysym.Sym {
				case sdl.K_ESCAPE:
					game.running = false
					break
				case sdl.K_SPACE:
					game.colored = true
				}
			} else if e.Type == sdl.KEYUP {
				switch e.Keysym.Sym {
				case sdl.K_SPACE:
					game.colored = false
				}
			}
		}
	}
}

func display(game *Game, renderer *sdl.Renderer) {
	if game.colored {
		renderer.SetDrawColor(161, 0, 255, 255)
	} else {
		renderer.SetDrawColor(211, 255, 142, 255)
	}
	renderer.Clear()
	renderer.Present()
}