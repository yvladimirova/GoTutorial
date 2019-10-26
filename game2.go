package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"io/ioutil"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	//"worldiety.de/training03/assets"
)

type Game struct {
	ShipImg     *ebiten.Image
	shipImgOpts *ebiten.DrawImageOptions
	x           float64
	y           float64
}

func (g *Game) setup() error {
	g.x = 2.0
	g.y = 3.0

	//img, _, err := image.Decode(bytes.NewReader(assets.Ship_png))
	dat, err := ioutil.ReadFile("1.png")
	img, _, err := image.Decode(bytes.NewReader(dat))

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

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.y += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.x -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.x += 2
	}

	opts := &ebiten.DrawImageOptions{}
	//x, y := ebiten.CursorPosition()
	//opts.GeoM.Reset()
	opts.GeoM.Translate(g.x, g.y)
	screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff}) //Hintergrundfarbe
	ebitenutil.DebugPrint(screen, "2D EbitenGame")
	must(screen.DrawImage(g.ShipImg, opts))
	return nil
}
