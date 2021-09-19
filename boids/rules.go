package boids

type Rule interface {
	Apply(flock Flock, index int) Vector
}

type Ruler []Rule

func NewDefaultRuler() Ruler {
	rules := make([]Rule, 0)
	ar := AlignmentRule{
		neighborDistance: 50,
		factor:           10,
	}
	cr := CohesionRule{
		factor: .01,
	}
	sr := SeparationRule{
		neighborDistance: 50,
		factor:           .01,
	}
	rules = append(rules, ar, cr, sr)
	return rules
}

type AlignmentRule struct {
	neighborDistance float64
	factor           float64
}

func (a AlignmentRule) Apply(flock Flock, index int) Vector {
	v := Vector{}
	nbBoid := 0
	for i, boid := range flock {
		if i == index {
			continue
		}
		if flock[index].IsNeighbor(*boid, a.neighborDistance) {
			v.add(boid.velocity)
			nbBoid++
		}
	}
	if nbBoid == 0 {
		return v
	}
	v.X = v.X / float64(nbBoid)
	v.Y = v.Y / float64(nbBoid)

	origV := flock[index].velocity
	v.X = (v.X - origV.X) * a.factor / 100
	v.Y = (v.Y - origV.Y) * a.factor / 100
	return v
}

type CohesionRule struct {
	factor float64
}

func (c CohesionRule) Apply(flock Flock, index int) Vector {
	v := Vector{}
	nbBoid := 0
	for i, boid := range flock {
		if i == index {
			continue
		}
		v.add(boid.position)
		nbBoid++
	}
	if nbBoid == 0 {
		return v
	}
	v.X = v.X / float64(nbBoid)
	v.Y = v.Y / float64(nbBoid)

	origV := flock[index].position
	v.X = (v.X - origV.X) * c.factor / 100
	v.Y = (v.Y - origV.Y) * c.factor / 100
	return v
}

type SeparationRule struct {
	neighborDistance float64
	factor           float64
}

func (s SeparationRule) Apply(flock Flock, index int) Vector {
	v := Vector{}
	origP := flock[index].position
	for i, boid := range flock {
		if i == index {
			continue
		}
		if flock[index].IsNeighbor(*boid, s.neighborDistance) {
			v.sub(Vector{
				X: boid.position.X - origP.X,
				Y: boid.position.Y - origP.Y,
			})
		}
	}
	v.X *= s.factor / 100
	v.Y *= s.factor / 100
	return v
}
