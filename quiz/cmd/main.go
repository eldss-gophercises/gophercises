package main

import (
	"fmt"
	"log"
	"os"

	quizlib "github.com/eldss/gophercises/quiz/lib"
)

func main() {
	file, err := os.Open("../problems.csv")
	if err != nil {
		log.Fatalln("Problem opening file:", err)
	}

	list, err := quizlib.ReadCsvFromReader(file)
	if err != nil {
		log.Fatalln("Problem parsing csv:", err)
	}
	file.Close()

	fmt.Println(list)
}
