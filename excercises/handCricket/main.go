package main

import (
	"fmt"

	"cricket.com/services"
)

func main() {
	cricket := services.Cricket{}
	cricket.Round1 = services.Round{}
	cricket.Round1.Batter.Name = "A"
	cricket.Round1.Bowler.Name = "B"

	winner := cricket.StartCricket()
	if winner == nil {
		fmt.Println("Tie!!!")
	} else {
		fmt.Printf("Winner is %s \n", winner.Name)
	}

}
