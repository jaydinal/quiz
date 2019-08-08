package quizecore

import (
	"strings"
)

// SimpleQuestion : struct type defined for Simple Question
type SimpleQuestion struct {
	question string
	answer   string
}

//Question : Question inteface method for Question
func (q SimpleQuestion) Question() string {
	return q.question
}

//Validate : Question inteface method for Validation
func (q SimpleQuestion) Validate(ans string) bool {
	return strings.EqualFold(ans, q.answer)
}

//Answer : Question inteface method for Answer
func (q SimpleQuestion) Answer() string {
	return q.answer
}

//Score : Question inteface method for Answer
func (q SimpleQuestion) Score() int {
	return 10
}

//NewSimpleQuestion : create new simple question
func NewSimpleQuestion(q string, a string) *SimpleQuestion {
	return &SimpleQuestion{q, a}
}
