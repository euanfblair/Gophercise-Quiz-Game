package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/euanfblair/Gophercise-Quiz-Game/quiz"
	"github.com/euanfblair/Gophercise-Quiz-Game/reader"
)

func main() {

	score := 0

	// Define flags
	fileName := flag.String("file", "./problems.csv", "the quiz file csv")
	timerLength := flag.String("timer", "30", "the timer length in seconds")

	// Parse the flags
	flag.Parse()

	// Validate and convert timer length
	timeInt, err := strconv.Atoi(*timerLength)
	if err != nil {
		fmt.Printf("timer flag should be a whole number\n")
		log.Fatal(err)
	}

	questions_and_answers, err := reader.ReadCsvToSlice(*fileName)
	questions := len(questions_and_answers)
	if err != nil {
		log.Fatal(err)
	}
	// Channel to signal when the quiz is done
	done := make(chan bool)

	// Run the quiz in a goroutine
	go func() {
		quiz.RunQuiz(questions_and_answers, &score)
		done <- true
	}()

	// Run the timer concurrently
	go func() {
		time.Sleep(time.Duration(timeInt) * time.Second)
		done <- false
	}()
	
	//Clear terminal 
	fmt.Print("\033[H\033[2J")

	// Wait for either the quiz to finish or the timer to expire
	timeUp := <-done

	switch {
	case !timeUp:
		fmt.Printf("You ran out of time. You scored %d/%d\n", score, questions)
	case timeUp:
		fmt.Printf("You scored %d/%d\n", score, questions)
	}

}
