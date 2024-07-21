package quiz

import (
	"fmt"
	"time"
)

func RunQuiz(questions_and_answers [][]string, correct *int) int {

	var userAnswer string
	score := 0
	counter := 0

	for i, qa := range questions_and_answers {

		question := qa[0]
		answer := qa[1]
		counter = i + 1

		fmt.Printf("%d. %s:\n", counter, question)
		fmt.Scan(&userAnswer)
		if userAnswer == answer {
			*correct += 1
			fmt.Println("Correct! Well Done")
			continue
		}
		fmt.Println("Incorrect! Better Luck Next Time")
	}

	return score
}

func GameTimer(gameTime int) bool {

	timer := time.NewTimer(time.Duration(gameTime) * time.Second)
	<-timer.C
	return true
}
