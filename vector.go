package main

type Vector struct {
	X, Y float64
}

func (v *Vector) add(v2 Vector) {
	v.X += v2.X
	v.Y += v2.Y
}

func (v *Vector) sub(v2 Vector) {
	v.X -= v2.X
	v.Y -= v2.Y
}
