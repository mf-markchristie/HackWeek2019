package main

import (
	"flag"
	"log"
	"os"
	"breakout"
	"runtime/pprof"
	"github.com/hajimehoshi/ebiten"
)


func main() {
	game := &breakout.Game{}
	update := game.Update
	if err := ebiten.Run(update, breakout.ScreenWidth, breakout.ScreenHeight, 2, "Breakout"); err != nil {
		log.Fatal(err)
	}
}

