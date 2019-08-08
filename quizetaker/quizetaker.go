package main

import (
	"fmt"

	"github.com/jaydinal/quizecore"
)

func main() {
	var quize []quizecore.Question
	quize = append(quize, quizecore.NewSimpleQuestion("What colour is the sky? \n", "none"))
	quize = append(quize, quizecore.NewSimpleQuestion("What is the capital of India \n", "i"))
	quize = append(quize, quizecore.NewChoiceQuestion("Who is Amitabh?", "2", "Prist", "Actor", "Singer", "Busnessman"))
	score := 0
	for _, q := range quize {
		fmt.Printf(q.Question())
		var ans string
		fmt.Print("Your answer \n")
		fmt.Scanln(&ans)
		if q.Validate(ans) {
			mi, status := q.(quizecore.ScoreMultiplier)
			if status {
				score += mi.Score()
			} else {
				score++
			}

		} else {
			fmt.Printf("Wrong! The correct answer is %v \n", q.Answer())
		}
	}
	fmt.Printf("Your score is %v \n", score)
}
