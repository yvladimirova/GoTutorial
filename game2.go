package main

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image"
	"image/color"
	"github.com/hajimehoshi/ebiten"
	"time"
	"worldiety.de/training03/assets"
)

type Game struct {
ShipImg *ebiten.Image
shipImgOpts *ebiten.DrawImageOptions
x float64
y float64
}

var newTime, oldTime float64

func (g *Game) setup() error {
	g.x = 2.0
	g.y = 3.0

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

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.y += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.x -=2
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.x += 2
	}

	opts := &ebiten.DrawImageOptions{}
	//x, y := ebiten.CursorPosition()
	//opts.GeoM.Reset()
	opts.GeoM.Translate(g.x,g.y)
	screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff}) //Hintergrundfarbe
	ebitenutil.DebugPrint(screen, "2D EbitenGame")
	must(screen.DrawImage(g.ShipImg, opts))
	return nil
}

func duration() float64  {
	newTime = float64(time.Now().UnixNano())
	deltaTime := (newTime - oldTime) / float64(time.Millisecond)
	oldTime = newTime
	return deltaTime
}
