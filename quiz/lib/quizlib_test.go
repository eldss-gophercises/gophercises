package quizlib

import (
	"fmt"
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

func TestShuffleQuestions(t *testing.T) {
	in := `1,1
2,2
3,3
4,4
5,5
6,6
7,7
8,8
9,9
10,10
11,11
12,12
13,13
14,14
15,15
`
	reader := strings.NewReader(in)
	quiz, err := NewQuizFromCsvReader(reader)
	if err != nil {
		t.Errorf("Error in function: %v", err)
	}

	for i := 0; i < 10; i++ {
		// Get string representation of list
		list := fmt.Sprintf("%v", quiz.questions)
		// Shuffle and get new list
		quiz.ShuffleQuestions()
		newList := fmt.Sprintf("%v", quiz.questions)

		// Lists should be different every time (except on very rare occasions)
		if list == newList {
			t.Errorf("Expected a shuffled list. Before: %v - After: %v", list, newList)
		}
	}
}
