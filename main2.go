package main

import (

	"github.com/hajimehoshi/ebiten"
	"log"
	"worldiety.de/training03/assets"

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
	if err := ebiten.Run(mygame.render,  640, 480, 1.0, "mygame"); err != nil {
		log.Fatal(err)
	}
}
