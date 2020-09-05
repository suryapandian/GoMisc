package services

import "math/rand"

type Player struct {
	Name    string
	Gesture int
	Score   int
}

const maxThrow = 6
const invalidThrow = 5

func (p *Player) getGesture() (g int) {
	g = rand.Intn(maxThrow)
	if g == invalidThrow {
		g = p.getGesture()
	}
	return
}

func (p *Player) updateScore(gesture int) {
	p.Score = p.Score + gesture
}
