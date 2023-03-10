# Ebintengine - Boids Simulation

This boids simulation is written in Go and utilizes the game library Ebintengine.
Check out the online version at https://rangzen.github.io/go-ebiten-boids/.

For more information on Ebintengine, please refer to the documentation at https://ebitengine.org/.
You can also check out author Hajime Hoshi's work on GitHub at https://github.com/hajimehoshi.

## Boids

Boids is an artificial life simulation program developed by Craig Reynolds in 1986.
The aim of the simulation was to replicate the flocking behavior of birds and other animals.
Instead of controlling the interactions of an entire flock, however,
the Boids simulation only specifies the behavior of each individual bird.
The Boids algorithm works by having each "boid" steer itself based on
rules of avoidance, alignment, and cohesion.
There are various demonstrations and implementations of the Boids algorithm available online.

See https://en.wikipedia.org/wiki/Boids for more information.