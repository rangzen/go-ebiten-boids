package main

import (
	"image/color"
	"math"
	"math/rand"
)

type Boid struct {
	position Vector
	velocity Vector
	color    color.RGBA
}

func NewBoid() *Boid {
	b := Boid{
		position: Vector{
			X: rand.Float64() * float64(originalSize),
			Y: rand.Float64() * float64(originalSize),
		},
		velocity: Vector{
			X: randomVelocity(),
			Y: randomVelocity(),
		},
		color: color.RGBA{
			R: uint8(int8(rand.Float64() * 0xff)),
			G: uint8(int8(rand.Float64() * 0xff)),
			B: uint8(int8(rand.Float64() * 0xff)),
			A: 0xff,
		},
	}
	return &b
}

func randomVelocity() float64 {
	return rand.Float64()*maxSpeed - maxSpeed/2
}

func (b *Boid) IsNeighbor(b2 Boid, distance float64) bool {
	// return math.Sqrt(math.Pow(b.position.X-b2.position.X, 2)+math.Pow(b.position.Y-b2.position.Y, 2)) < distance
	return math.Abs(b.position.X-b2.position.X) < distance && math.Abs(b.position.Y-b2.position.Y) < distance
}

func (b *Boid) Update(width, height int) {
	b.position.X = overLimit(b.position.X+b.velocity.X, width)
	b.position.Y = overLimit(b.position.Y+b.velocity.Y, height)
}

func (b *Boid) SpeedLimit(min, max float64) {
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

func overLimit(position float64, limit int) float64 {
	if position > float64(limit) {
		return position - float64(limit)
	} else if position < 0 {
		return position + float64(limit)
	}
	return position
}
