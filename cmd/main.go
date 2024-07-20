package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/euanfblair/Gophercise-Quiz-Game/quiz"
	"github.com/euanfblair/Gophercise-Quiz-Game/reader"
)

func main() {

	timerLength := reader.FileFlag("timer", "30", "the timer length in seconds")

	timeInt, err := strconv.Atoi(timerLength)
	if err != nil {
		fmt.Printf("timer flag should be a whole number")
		log.Fatal(err)
	}

	timer := time.NewTimer(time.Duration(timeInt) * time.Second)

	go func() {

		fileName := reader.FileFlag("file", "./problems.csv", "the quiz file csv")

		questions_and_answers, err := reader.ReadCsvToSlice(fileName)
		if err != nil {
			log.Fatal(err)
		}
		score, total := quiz.RunQuiz(questions_and_answers)
		fmt.Print("\033[H\033[2J")
		fmt.Printf("You scored %d/%d\n", score, total)
		os.Exit(0)
	}()

	<-timer.C
	fmt.Println("You ran out of time, better luck next time")

}
