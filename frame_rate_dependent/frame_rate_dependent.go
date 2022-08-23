package frame_rate_dependent

import (
	"fmt"
	"github.com/ungame/go-deltatime/core"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

func Run(fps float64) {
	fmt.Printf("TARGET: %.0ffps\n\n", fps)

	var (
		running    = true
		frameStart time.Time
		frameTime  time.Duration
		frameRate  = 1000 / fps
		deltaTime  float64
	)

	var (
		object = sdl.FRect{X: 0, Y: 50, W: 50, H: 50}
		speed  = float32(10)
	)

	var (
		iterations          int
		iterationsPerSecond = 0
		timerPerSecond      = time.Now()
	)

	var (
		startedAt time.Time
		start     bool
	)

	for running {
		frameStart = time.Now()

		if !core.Listen() || object.X+object.W >= core.SCREEN_WIDTH {
			running = false
		}

		if core.KeyPressed(sdl.SCANCODE_SPACE) {
			start = true
			startedAt = time.Now()
		}

		if start {
			object.X += speed
		}

		core.Draw(func(renderer *sdl.Renderer) {
			_ = renderer.SetDrawColor(189, 189, 189, sdl.ALPHA_TRANSPARENT)
			_ = renderer.FillRectF(&object)
		})

		frameTime = time.Since(frameStart)
		deltaTime = float64(frameTime.Milliseconds())

		if deltaTime < frameRate {
			sdl.Delay(uint32(frameRate - deltaTime))
		}

		iterations++
		if time.Since(timerPerSecond) >= time.Second {
			iterationsPerSecond++
			timerPerSecond = time.Now()
			fmt.Printf("\rFPS: %d", iterations/iterationsPerSecond)
		}
	}

	if start {
		fmt.Println("")
		fmt.Println("Elapsed:", time.Since(startedAt).String())
	}
}
