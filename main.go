package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	"log"
	"math/rand"
	"os"
	"time"
)

const exitCodeNormal = 0

const (
	originalSize      = 480 // original size of the window
	startingFlockSize = 500 // how many boid created at start
	tickPeriod        = 1   // how many ticks between update
	minSpeed          = 2.  // max speed
	maxSpeed          = 4.  // max speed
	vectorRatio       = 7   // ratio for vector drawing
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (g *Game) Update() error {
	// Keyboard input
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.AddBoid()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.DelBoid()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyV) {
		g.resetFlockVelocity()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		os.Exit(exitCodeNormal)
	}

	// Flock update
	g.timer++
	if g.timer%g.moveTime == 0 {
		g.updateFlock()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Background
	screen.Fill(color.RGBA{R: 0x64, G: 0xb4, B: 0xb4, A: 0xff})

	// Flock
	for _, b := range g.flock {
		ebitenutil.DrawLine(screen,
			b.position.X,
			b.position.Y,
			b.position.X+b.velocity.X*vectorRatio,
			b.position.Y+b.velocity.Y*vectorRatio,
			b.color,
		)
	}

	// Debug
	ebitenutil.DebugPrint(screen, fmt.Sprintf(" FPS: %0.2f\n Flock size: %d", ebiten.CurrentFPS(), len(g.flock)))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	g.width = outsideWidth
	g.height = outsideHeight
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(originalSize, originalSize)
	ebiten.SetWindowTitle("Boid Flock in Go/Ebiten")
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
