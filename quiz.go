package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
)

func main() {
	fileName := flag.String("csv", "questions.csv", "the filename for the csv with the question,answer")
	flag.Parse()

	questionFile, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	questionReader := csv.NewReader(questionFile)
	for {
		// Read from reader
		// Check for EOF
		// Print question
		// store answer to each question
	}
}
