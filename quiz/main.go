package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Quiz struct {
	question string
	answer   string
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "File name of the quiz problems")
	limit := flag.Int("limit", 10, "Time limit(in seconds) for quiz")
	flag.Parse()

	data, err := ParseCSV(*csvFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	quizData := CreateQuizList(data)

	var correctAns int
	limitAlert := time.After(time.Second * time.Duration(*limit))
	go PlayQuiz(quizData, &correctAns)
	<-limitAlert

	fmt.Printf("\nYou scored %d out of %d", correctAns, len(quizData))
}

func ParseCSV(filepath string) ([][]string, error) {
	var data [][]string
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		return data, err
	}

	csvReader := csv.NewReader(file)
	data, err = csvReader.ReadAll()
	if err != nil {
		return data, err
	}
	return data, nil
}

func CreateQuizList(data [][]string) []Quiz {
	quiz := make([]Quiz, len(data))
	for i, line := range data {
		quiz[i] = Quiz{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return quiz
}

func PlayQuiz(quizData []Quiz, correctAns *int) {

	for i, q := range quizData {
		var ans string
		fmt.Printf("Problem #%d: %s = ", i+1, q.question)
		_, err := fmt.Scanln(&ans)
		if err != nil {
			continue
		} else if ans == q.answer {
			*correctAns++
		}
	}
}
