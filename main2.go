package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
)

var shipImg *ebiten.Image
var newTime, oldTime float64

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	mygame := Game{}
	mygame.setup()
	if err := ebiten.Run(mygame.render, 640, 480, 1.0, "mygame"); err != nil {
		log.Fatal(err)
	}
}

func duration() float64 {
	newTime = float64(time.Now().UnixNano())
	deltaTime := (newTime - oldTime) / float64(time.Millisecond)
	oldTime = newTime
	return deltaTime
}
