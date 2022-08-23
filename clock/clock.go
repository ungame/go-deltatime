package clock

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

var (
	prevTime  time.Time
	frameRate float64
	deltaTime float64
)

func init() {
	prevTime = time.Now()
}

func Tick(fps float64) float64 {
	frameRate = 1000 / fps
	deltaTime = float64(time.Since(prevTime).Milliseconds())
	prevTime = time.Now()
	if frameRate > deltaTime {
		sdl.Delay(uint32(frameRate - deltaTime))
	}
	return deltaTime
}
