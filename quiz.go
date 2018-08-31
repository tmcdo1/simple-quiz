package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fileName := flag.String("csv", "questions.csv", "the filename for the csv with the question,answer")
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

	fmt.Printf("You go %d out of %d questions correctly!\n", numCorrect, numQuestions)
}
