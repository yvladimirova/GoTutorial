package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"log"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"worldiety.de/training03/assets"
)

type Game struct {
ShipImg *ebiten.Image
shipImgOpts *ebiten.DrawImageOptions
x,y float64
}

func (g *Game) setup() error {
	img, _, err := image.Decode(bytes.NewReader(assets.Ship_png))

	if err != nil {
		return fmt.Errorf("unable to decode ship %w", err)
	}

	g.ShipImg, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	if err != nil {
		return fmt.Errorf("unable to load texture %w", err)
	}

	return nil
}

func (g *Game) render(screen *ebiten.Image) error {
	opts := &ebiten.DrawImageOptions{}
	x, y := ebiten.CursorPosition()
	//opts.GeoM.Reset()
	opts.GeoM.Translate(float64(x),float64(y))
	screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff}) //Hintergrundfarbe
	ebitenutil.DebugPrint(screen, "2D EbitenGame")
	must(screen.DrawImage(g.ShipImg, opts))
	return nil
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	g := Game{}
	g.setup()
	if err := ebiten.Run(g.render,  640, 480, 1.0, "mygame"); err != nil {
		log.Fatal(err)
	}
}
