package quizlib

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

// Quiz holds information about a single quiz.
type Quiz struct {
	questions []qaPair
}

// qaPair contains a single question and answer for a quiz
type qaPair struct {
	question string
	answer   string
}

// NewQuizFromCsvReader takes an io.Reader pointing to csv formatted data and
// parses it into a slice of QAPairs. If the csv has headers, they will be
// included as the first item in the returned slice.
func NewQuizFromCsvReader(r io.Reader) (Quiz, error) {
	var pairs = make([]qaPair, 0)

	csvR := csv.NewReader(r)
	csvR.FieldsPerRecord = 2
	for {
		record, err := csvR.Read()
		// End of file
		if err == io.EOF {
			break
		}
		// Some error occurred
		if err != nil {
			return Quiz{}, err
		}
		// Extract and collect question/answer
		pair := qaPair{record[0], record[1]}
		pairs = append(pairs, pair)
	}

	var quiz = Quiz{pairs}

	return quiz, nil
}

// GetQA returns a question and answer given a question number (0 indexed).
// Returns an error if the question number is larger than the list of questions.
func (q *Quiz) GetQA(qNumber int) (string, string, error) {
	if qNumber >= 0 && qNumber < q.NumQuestions() {
		qa := q.questions[qNumber]
		return qa.question, qa.answer, nil
	}

	// Error case
	errStr := fmt.Sprintf("Question number %d does not exist. %d questions in this quiz.",
		qNumber, q.NumQuestions())
	err := errors.New(errStr)
	return "", "", err
}

// ShuffleQuestions randomly changes the order of questions in the Quiz
func (q *Quiz) ShuffleQuestions() {
	rand.Seed(time.Now().UnixNano())
	length := q.NumQuestions()
	for i := 0; i < length; i++ {
		newI := rand.Intn(length)
		q.questions[i], q.questions[newI] = q.questions[newI], q.questions[i]
	}
}

// NumQuestions gives the length of the quiz
func (q *Quiz) NumQuestions() int {
	return len(q.questions)
}
