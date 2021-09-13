package main

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
	rules = append(rules, ar)
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
