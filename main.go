package main

import (
	"flag"
	"github.com/ungame/go-deltatime/core"
	"github.com/ungame/go-deltatime/frame_rate_dependent"
	"github.com/ungame/go-deltatime/frame_rate_independent"
)

var (
	fps         float64
	independent bool
)

func init() {
	flag.Float64Var(&fps, "fps", 60, "set fps to run")
	flag.BoolVar(&independent, "i", true, "set true to run frame rate independent")
	flag.Parse()
}

func main() {
	defer core.Quit()
	if independent {
		frame_rate_independent.Run(fps)
	} else {
		frame_rate_dependent.Run(fps)
	}
}
