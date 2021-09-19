package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rangzen/go-ebiten-boids/boids"
	"log"
)

func main() {
	ebiten.SetWindowSize(boids.OriginalSize, boids.OriginalSize)
	ebiten.SetWindowTitle("Boid Flock in Go/Ebiten")
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(boids.NewGame()); err != nil {
		log.Fatal(err)
	}
}
