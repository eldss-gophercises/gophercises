package quizlib

import (
	"strings"
	"testing"
)

func TestNewQuiz(t *testing.T) {
	in := `water is wet,true
sun is bright,true
chairs are for eating,false	
`
	reader := strings.NewReader(in)
	quiz, err := NewQuizFromCsvReader(reader)
	if err != nil {
		t.Errorf("Error in function: %v", err)
	}

	if quiz.NumQuestions() != 3 {
		t.Errorf("Length of questions slice should be 3. Questions: %v",
			quiz.questions)
	}
}

func TestGetQA(t *testing.T) {
	in := `water is wet,true
sun is bright,true
`
	reader := strings.NewReader(in)
	quiz, err := NewQuizFromCsvReader(reader)
	if err != nil {
		t.Errorf("Error in function: %v", err)
	}

	q, a, err := quiz.GetQA(0)
	if q != "water is wet" || a != "true" {
		t.Errorf("Got q: %v, a: %v", q, a)
	}

	q, a, err = quiz.GetQA(1)
	if q != "sun is bright" || a != "true" {
		t.Errorf("Got q: %v, a: %v", q, a)
	}

	// Error case
	q, a, err = quiz.GetQA(2)
	if err == nil {
		t.Error("Expected an error. Asked for an non-existant question.")
	}
}
