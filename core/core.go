package core

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

var (
	window   *sdl.Window
	renderer *sdl.Renderer
	keyboard []uint8
)

const (
	WINDOW_TITLE  = "go-deltatime"
	SCREEN_WIDTH  = 1200
	SCREEN_HEIGHT = 720
)

func init() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		log.Fatalln("unable to init sdl:", err.Error())
	}
	window, err = sdl.CreateWindow(
		WINDOW_TITLE,
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		SCREEN_WIDTH,
		SCREEN_HEIGHT,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		log.Fatalln("unable to create window:", err.Error())
	}
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatalln("unable to create renderer:", err.Error())
	}
	keyboard = sdl.GetKeyboardState()
}

func Listen() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.GetType() {
		case sdl.QUIT:
			return false
		case sdl.KEYDOWN, sdl.KEYUP:
			keyboard = sdl.GetKeyboardState()
		}
	}
	return true
}

func KeyPressed(key sdl.Scancode) bool {
	return keyboard[key] == 1
}

func Draw(draw func(*sdl.Renderer)) {
	err := renderer.SetDrawColor(0, 0, 0, sdl.ALPHA_TRANSPARENT)
	if err != nil {
		log.Println("unable to draw background:", err.Error())
		return
	}
	err = renderer.Clear()
	if err != nil {
		log.Println("unable to draw background:", err.Error())
		return
	}
	draw(renderer)
	renderer.Present()
}

func Quit() {
	destroyRenderer()
	destroyWindow()
	sdl.Quit()
}

func destroyRenderer() {
	if renderer != nil {
		err := renderer.Destroy()
		if err != nil {
			log.Println("unable to destroy renderer:", err.Error())
		}
	}
}

func destroyWindow() {
	if window != nil {
		err := window.Destroy()
		if err != nil {
			log.Println("unable to destroy window:", err.Error())
		}
	}
}
