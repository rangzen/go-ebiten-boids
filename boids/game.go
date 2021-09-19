package boids

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	"math/rand"
	"os"
	"time"
)

const exitCodeNormal = 0

const (
	OriginalSize      = 480 // original size of the window
	startingFlockSize = 108 // how many boid created at start
	tickPeriod        = 1   // how many ticks between update
	minSpeed          = 2.  // max speed
	maxSpeed          = 4.  // max speed
	vectorRatio       = 7   // ratio for vector drawing
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Game struct {
	flock         Flock
	ruler         Ruler
	timer         int
	moveTime      int
	width, height int
}

type Flock []*Boid

// Update is part of ebiten.Game implementation
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

// Draw is part of ebiten.Game implementation
func (g *Game) Draw(screen *ebiten.Image) {
	// Background
	screen.Fill(color.RGBA{})

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

// Layout is part of ebiten.Game implementation
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	g.width = outsideWidth
	g.height = outsideHeight
	return outsideWidth, outsideHeight
}

func NewGame() *Game {
	flock := make([]*Boid, 0, startingFlockSize)
	for i := 0; i < startingFlockSize; i++ {
		flock = append(flock, NewBoid())
	}
	g := Game{
		flock:    flock,
		ruler:    NewDefaultRuler(),
		moveTime: tickPeriod,
	}
	return &g
}

func (g *Game) AddBoid() {
	g.flock = append(g.flock, NewBoid())
}

func (g *Game) DelBoid() {
	if len(g.flock) > 1 {
		g.flock = g.flock[:len(g.flock)-1]
	}
}

func (g *Game) updateFlock() {
	nextFlock := make(Flock, len(g.flock))
	copy(nextFlock, g.flock)

	for i, b := range nextFlock {
		for _, rule := range g.ruler {
			vector := rule.Apply(g.flock, i)
			b.velocity.add(vector)
		}
		b.Update(g.width, g.height)
	}

	g.flock = nextFlock
}

func (g *Game) resetFlockVelocity() {
	for _, b := range g.flock {
		b.velocity.X = randomVelocity()
		b.velocity.Y = randomVelocity()
	}
}
