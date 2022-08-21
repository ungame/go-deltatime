package main

import (
	"flag"
	"github.com/ungame/go-deltatime/core"
	"github.com/ungame/go-deltatime/independent_frame_rate"
)

var fps float64

func init() {
	flag.Float64Var(&fps, "fps", 60, "set fps to run")
	flag.Parse()
}

func main() {
	defer core.Quit()
	independent_frame_rate.Run(fps)
}
