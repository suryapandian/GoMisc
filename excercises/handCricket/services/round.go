package services

import "fmt"

type Round struct {
	Batter Player
	Bowler Player
}

const numOfThrows = 6

func (r *Round) startRound() *Round {
	for i := 0; i <= numOfThrows; i++ {
		r.Batter.Gesture = r.Batter.getGesture()
		r.Bowler.Gesture = r.Bowler.getGesture()
		fmt.Printf("%s throws %d, %s throws %d \n", r.Batter.Name, r.Batter.Gesture, r.Bowler.Name, r.Bowler.Gesture)

		if r.Batter.Gesture == r.Bowler.Gesture {
			return r
		}

		r.Batter.updateScore(r.Batter.Gesture)
		fmt.Printf("%s's Score is %d \n", r.Batter.Name, r.Batter.Score)
	}
	return r
}
