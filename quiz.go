package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	fileName := flag.String("csv", "questions.csv", "the filename for the csv with the question,answer")
	timeLimit := flag.Int("limit", -1, "Specify a time limit you would like to have for the quiz in seconds")
	flag.Parse()

	questionFile, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	questionReader := csv.NewReader(questionFile)

	records, err := questionReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	numQuestions := len(records)
	numCorrect := 0

	if *timeLimit > 0 {
		tTicker := time.NewTicker(time.Duration(int64(*timeLimit)) * time.Second)
		go func() {
			<-tTicker.C
			tTicker.Stop()
			end(numCorrect, numQuestions)
		}()
	}

	done := make(chan bool)
	go func() {
		for _, v := range records {
			answer := strings.ToLower(strings.TrimSpace(v[1]))
			fmt.Println(v[0])
			var res string
			fmt.Scanf("%s\n", &res)
			res = strings.ToLower(strings.TrimSpace(res))
			if res == answer {
				numCorrect++
			}
		}
		done <- true
	}()

	<-done
	end(numCorrect, numQuestions)
}

func end(correct int, total int) {
	fmt.Printf("You got %d out of %d questions correct!\n", correct, total)
	os.Exit(0)
}
