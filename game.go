package main

type Game struct {
	flock         Flock
	ruler         Ruler
	timer         int
	moveTime      int
	width, height int
}

type Flock []*Boid

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
		b.SpeedLimit(minSpeed, maxSpeed)
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
