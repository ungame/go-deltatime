package frame_rate_independent

import (
	"fmt"
	"github.com/ungame/go-deltatime/core"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

func Run(fps float64) {
	fmt.Printf("TARGET: %.0f FPS\n\n", fps)
	// frame settings
	var (
		frameRate        = 1000 / fps // FPS converted to milliseconds
		frameStart       time.Time
		prevTime         time.Time
		deltaTime        = float64(0)
		frameTime        time.Duration
		frames           int
		counterPerSecond int
		timerPerSecond   time.Time
		startedAt        time.Time
	)

	// Object settings
	var (
		velocity = float64(300) // move 300 frames per second independent of frame rate.
		object   = sdl.Rect{X: 0, Y: 100, W: 50, H: 50}
		position = sdl.FPoint{}
	)

	// Control settings
	var (
		start   bool // press SPACE to start!
		running = true
	)

	timerPerSecond = time.Now()
	prevTime = time.Now()

	for running {
		frameStart = time.Now()

		if !core.Listen() || object.X+object.W >= core.SCREEN_WIDTH {
			running = false
		}

		if core.KeyPressed(sdl.SCANCODE_SPACE) && !start {
			start = true
			startedAt = time.Now()
		}

		if start {
			position.X += float32(velocity * deltaTime)
		}

		object.X = int32(position.X)

		core.Draw(func(renderer *sdl.Renderer) {
			_ = renderer.SetDrawColor(255, 255, 255, sdl.ALPHA_TRANSPARENT)
			_ = renderer.FillRect(&object)
		})

		frameTime = frameStart.Sub(prevTime)
		prevTime = frameStart
		deltaTime = frameTime.Seconds()

		frameTimeNoDelay := time.Since(frameStart)
		deltaTimeNoDelay := float64(frameTimeNoDelay.Milliseconds())
		if deltaTimeNoDelay < frameRate {
			sdl.Delay(uint32(frameRate - deltaTimeNoDelay))
		}

		frames++
		if time.Since(timerPerSecond) >= time.Second {
			counterPerSecond++
			timerPerSecond = time.Now()
			fmt.Printf("\rFPS=%d, FrameTime=%s", frames/counterPerSecond, frameTime.String())
		}
	}

	if start {
		fmt.Println("")
		fmt.Println("Elapsed:", time.Since(startedAt).String())
	}
}
