package main

import (
	"fmt"
	"github.com/ungame/go-deltatime/clock"
	"github.com/ungame/go-deltatime/core"
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"time"
)

func main() {

	var (
		object   = sdl.FRect{X: 0, Y: 50, W: 50, H: 50}
		position = sdl.FPoint{}
		speed    = float32(300)
	)

	var (
		iterations          int
		iterationsPerSecond = 0
		timerPerSecond      = time.Now()
	)

	var (
		startedAt time.Time
		start     bool
		running   = true
	)

	for running {
		dt := clock.Tick(60)

		if !core.Listen() || object.X+object.W >= core.SCREEN_WIDTH {
			running = false
		}

		if core.KeyPressed(sdl.SCANCODE_SPACE) {
			start = true
			startedAt = time.Now()
		}

		if start {
			sec := NewFloat(dt / 1000)
			position.X += speed * sec.Float32()
			object.X = round(position.X)
		}

		core.Draw(func(renderer *sdl.Renderer) {
			_ = renderer.SetDrawColor(189, 189, 189, sdl.ALPHA_TRANSPARENT)
			_ = renderer.FillRectF(&object)
		})

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

func round(f float32) float32 {
	return float32(math.Round(float64(f)))
}

type GFloat interface {
	float32 | float64
}

type Float[F GFloat] struct {
	v F
}

func NewFloat[F GFloat](f F) *Float[F] {
	return &Float[F]{v: f}
}

func (f *Float[F]) Float32() float32 {
	return float32(f.v)
}

func (f *Float[F]) Float64() float64 {
	return float64(f.v)
}
