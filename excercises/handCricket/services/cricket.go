package services

import (
	"fmt"
)

type Cricket struct {
	Round1 Round
	Round2 Round
}

func (c *Cricket) StartCricket() *Player {
	c.Round1.startRound()
	fmt.Println("Score after round1", c.Round1.Batter.Score)

	c.Round2.Batter = c.Round1.Bowler
	c.Round2.Bowler = c.Round1.Batter

	c.Round2.startRound()
	fmt.Println("Score after round2", c.Round2.Batter.Score)
	if c.Round1.Batter.Score > c.Round2.Batter.Score {
		return &c.Round1.Batter
	}
	if c.Round2.Batter.Score > c.Round1.Batter.Score {
		return &c.Round2.Batter
	}
	return nil
}
