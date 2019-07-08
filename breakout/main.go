package main

import (
	"flag"
	"log"
	"os"
	"breakout"
	"runtime/pprof"
	"github.com/hajimehoshi/ebiten"
)

var cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}

	game := &breakout.Game{}
	update := game.Update
	if err := ebiten.Run(update, breakout.ScreenWidth, breakout.ScreenHeight, 2, "Breakout"); err != nil {
		log.Fatal(err)
	}
}

