package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/img"
)

type Game struct {
	running bool
	score int
	colored bool
}
type Img struct {
	texture *sdl.Texture
	width int32
	height int32
}

var testImg = Img {}

func main() {
	game := Game {running: true}

	// Initialize SDL
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Printf("Could not initialize SDL2: %s\n", sdl.GetError())
		return
	}
	defer sdl.Quit()

	//Preload image library
	if img.Init(img.INIT_JPG) != img.INIT_JPG {
		fmt.Printf("Could not initialize image library: %s\n", img.GetError())
		return
	}
	defer img.Quit()

	// Create the window
	window, err := sdl.CreateWindow("SDL2", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 1024, 576, sdl.WINDOW_RESIZABLE)
	if err != nil {
		fmt.Printf("Could not create window: %s\n", sdl.GetError())
		return
	}
	defer window.Destroy()

	// Create renderer
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Printf("Could not create renderer: %s\n", sdl.GetError())
		return
	}
	defer renderer.Destroy()

	//Load image
	imgSurface, err := img.Load("adelitas.jpg")
	if err != nil {
		fmt.Printf("Failed to load image: %s\n", img.GetError())
	}
	testImg.width = imgSurface.W
	testImg.height = imgSurface.H

	//Create texture
	imgTexture, err := renderer.CreateTextureFromSurface(imgSurface)
	if err != nil {
		fmt.Printf("Could not create texture: %s\n", sdl.GetError())
	}
	testImg.texture = imgTexture
	imgSurface.Free()

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
	renderer.Copy(testImg.texture, nil, &sdl.Rect{50, 50, testImg.width/3, testImg.height/3})
	renderer.Present()
}