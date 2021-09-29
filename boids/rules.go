package boids

type Rule interface {
	Apply(flock Flock, index int) Vector
}

type Ruler struct {
	a   AlignmentRule
	c   CohesionRule
	s   SeparationRule
	att AttractorRule
}

func (r *Ruler) Apply(flock Flock, index int) Vector {
	v := Vector{}
	for _, rule := range []Rule{r.a, r.c, r.s} {
		vr := rule.Apply(flock, index)
		v.add(vr)
	}
	if r.att.attractor.IsAlive() {
		vr := r.att.Apply(flock, index)
		v.add(vr)
	}
	return v
}

func (r *Ruler) AttractorOn(x int, y int) {
	r.att.attractor.position.X = float64(x)
	r.att.attractor.position.Y = float64(y)
	r.att.attractor.ttl = attractorDefaultTtl
}

func NewDefaultRuler() Ruler {
	return Ruler{
		a: AlignmentRule{
			neighborDistance: 20,
			factor:           15,
		},
		c: CohesionRule{
			factor: .01,
		},
		s: SeparationRule{
			neighborDistance: 10,
			factor:           1,
		},
		att: AttractorRule{
			attractor: Attractor{
				position: Vector{},
				ttl:      0,
			},
			factor: .1,
		},
	}
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

type AttractorRule struct {
	attractor Attractor
	factor    float64
}

func (s *AttractorRule) Apply(flock Flock, index int) Vector {
	origP := flock[index].position
	v := Vector{
		X: s.attractor.position.X - origP.X,
		Y: s.attractor.position.Y - origP.Y,
	}
	v.X *= s.factor / 100
	v.Y *= s.factor / 100

	s.attractor.Lived()

	return v
}
