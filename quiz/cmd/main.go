package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	quizlib "github.com/eldss/gophercises/quiz/lib"
)

// Channel messages
const increment = 1
const exit = 0

// Flag vars
var csv string
var seconds uint
var minutes uint

// Setup Flags
func init() {
	var shorthand = " (shorthand)"

	// Csv file flag
	const (
		defaultCsv = ""
		usageCsv   = "A csv file with question,answer fields."
	)
	flag.StringVar(&csv, "csv", defaultCsv, usageCsv)
	flag.StringVar(&csv, "c", defaultCsv, usageCsv+shorthand)

	// Timer flag - seconds
	const (
		defaultSecs = 30
		usageSecs   = "Length of the test, in seconds."
	)
	flag.UintVar(&seconds, "secs", defaultSecs, usageSecs)
	flag.UintVar(&seconds, "s", defaultSecs, usageSecs+shorthand)

	// timer flag - minutes
	const (
		defaultMins = 0
		usageMins   = "Length of the test, in minutes."
	)
	flag.UintVar(&minutes, "mins", defaultMins, usageMins)
	flag.UintVar(&minutes, "m", defaultMins, usageMins+shorthand)
}

func main() {
	flag.Parse()

	// Validate csv filepath
	if csv == "" {
		fmt.Println("Must specify a filename:")
		flag.Usage()
		os.Exit(1)
	}

	file, err := os.Open(csv)
	if err != nil {
		log.Fatalln("Problem opening file:", err)
	}

	// Read file contents to Quiz
	quiz, err := quizlib.NewQuizFromCsvReader(file)
	if err != nil {
		log.Fatalln("Problem parsing csv:", err)
	}
	file.Close()

	// Set up and run quiz on a timer
	score := 0
	timeLimit := seconds + (minutes * 60)
	messages := make(chan int)
	go timer(timeLimit, messages)
	go runQuiz(quiz, os.Stdin, messages)

	// Listen for messages
	for {
		m := <-messages
		if m == exit {
			fmt.Println()
			break
		} else if m == increment {
			score++
		} else {
			log.Fatalln("An error occurred. Exiting quiz.")
		}
	}

	// Print stats
	fmt.Printf("You scored %d/%d\n", score, quiz.NumQuestions())
}

// Runs a quiz given the Quiz and a Reader to get user answers
func runQuiz(q quizlib.Quiz, r io.Reader, c chan int) {
	userIn := bufio.NewScanner(r)

	// let user begin when ready
	fmt.Println("Press Enter when ready.")
	userIn.Scan()

	for i := 0; i < q.NumQuestions(); i++ {
		// Get a question and answer
		question, answer, err := q.GetQA(i)
		if err != nil {
			log.Fatalln("Problem getting question, exiting quiz:", err)
		}

		// Ask the question and get an answer
		fmt.Println(question)
		userIn.Scan()
		fmt.Println("----------")

		// Clean answers
		answer = strings.ToLower(strings.TrimSpace(answer))
		userAnswer := strings.ToLower(strings.TrimSpace(userIn.Text()))

		// Validate answer
		if answer == userAnswer {
			c <- increment
		}
	}

	c <- exit
}

func timer(seconds uint, c chan int) {
	time.Sleep(time.Duration(seconds) * time.Second)
	c <- exit
}
