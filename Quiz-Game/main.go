package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

var path string = "./problems.csv"
var content [][]string
var totalQuestion, correctQuestions int

func main() {
	fileContent, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	reader := csv.NewReader(fileContent)

	content, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	totalQuestion = len(content)

	timer := time.After(30 * time.Second)

	done := make(chan bool)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)

		for i := range totalQuestion {
			relevantQuestion := content[i][0]
			relevantAnswer := content[i][1]

			fmt.Printf("Question number: %v \n", i)

			fmt.Println("What is the result of: ", relevantQuestion)

			scanner.Scan()

			if answer := scanner.Text(); answer == relevantAnswer {
				correctQuestions += 1
				fmt.Printf("Correct! you have answered %v / %v correct. \n", correctQuestions, totalQuestion)
			} else {
				fmt.Printf("Wrong answer... The correct answer was %s and your answer was %s \n", relevantAnswer, answer)
			}
		}
		done <- true
	}()

	select {
	case <-timer:
		fmt.Println("\n--- Time is up! ---")
	case <-done:
		fmt.Println("\n--- GZ you have completed the quiz! ---")
	}

	fmt.Println("You have completed the test")
}
