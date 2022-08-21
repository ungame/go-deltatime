package independent_frame_rate

import (
	"fmt"
	"github.com/ungame/go-deltatime/core"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

func Run(fps float64) {
	fmt.Printf("%.0fFPS\n\n", fps)
	// frame settings
	var (
		frameRate        = 1000 / fps // FPS converted to milliseconds
		startedAt        time.Time
		deltaTime        = float64(0)
		frameTime        time.Duration
		prevTime         = time.Now()
		frames           int
		counterPerSecond int
		timerPerSecond   time.Time
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

	for running {
		frameStart := time.Now()

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
		deltaTime = frameTime.Seconds()
		prevTime = frameStart

		frames++
		if time.Since(timerPerSecond) >= time.Second {
			counterPerSecond++
			timerPerSecond = time.Now()
			fmt.Printf("Frames=%d, PerSecond=%d\r", frames, frames/counterPerSecond)
		}

		sdl.Delay(uint32(frameRate))
	}

	if start {
		fmt.Println("")
		fmt.Println("Elapsed:", time.Since(startedAt).String())
	}
}
