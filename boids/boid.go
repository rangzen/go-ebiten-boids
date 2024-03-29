package boids

import (
	"image/color"
	"math"
	"math/rand"

	gc "github.com/gerow/go-color"
)

type Boid struct {
	position Vector
	velocity Vector
	color    color.RGBA
}

func NewBoid() *Boid {
	b := Boid{
		position: Vector{
			X: rand.Float64() * float64(OriginalSize),
			Y: rand.Float64() * float64(OriginalSize),
		},
		velocity: Vector{
			X: randomVelocity(),
			Y: randomVelocity(),
		},
		color: randomBrightColor(),
	}
	return &b
}

// Get a Bright Random Colour Python https://stackoverflow.com/a/43437435/337726
func randomBrightColor() color.RGBA {
	hsl := gc.HSL{
		H: rand.Float64(),
		S: .9,
		L: .6 + rand.Float64()/5,
	}
	rgb := hsl.ToRGB()
	return color.RGBA{
		R: uint8(rgb.R * 256),
		G: uint8(rgb.G * 256),
		B: uint8(rgb.B * 256),
		A: 0xff,
	}
}

func randomVelocity() float64 {
	return rand.Float64()*maxSpeed - maxSpeed/2
}

func (b *Boid) IsNeighbor(b2 Boid, distance float64) bool {
	// return math.Sqrt(math.Pow(b.position.X-b2.position.X, 2)+math.Pow(b.position.Y-b2.position.Y, 2)) < distance
	return math.Abs(b.position.X-b2.position.X) < distance && math.Abs(b.position.Y-b2.position.Y) < distance
}

func (b *Boid) Update(width int, height int) {
	b.checkSpeed(minSpeed, maxSpeed)
	b.position.X += b.velocity.X
	b.position.Y += b.velocity.Y
	b.checkPosition(width, height)
}

// checkPosition deals with window border
func (b *Boid) checkPosition(width, height int) {
	b.position.X = modulo(b.position.X, width)
	b.position.Y = modulo(b.position.Y, height)
}

// modulo keep a value in a modulo range
func modulo(v float64, m int) float64 {
	if v > float64(m) {
		return v - float64(m)
	} else if v < 0 {
		return v + float64(m)
	}
	return v
}

// checkSpeed keep speed in a range
func (b *Boid) checkSpeed(min, max float64) {
	d := math.Sqrt(math.Pow(b.velocity.X, 2) + math.Pow(b.velocity.Y, 2))
	if d < min {
		r := d / min
		b.velocity.X /= r
		b.velocity.Y /= r
	} else if d > max {
		r := d / max
		b.velocity.X /= r
		b.velocity.Y /= r
	}
}
