package quizecore

import "fmt"

// ChoiceQuestion : struct type defined for choice Question
type ChoiceQuestion struct {
	question string
	choices  []string
	answer   string
}

//Validate : Question inteface method for Validation
func (q ChoiceQuestion) Validate(ans string) bool {
	return ans == q.answer
}

//Answer : Question inteface method for Answer
func (q ChoiceQuestion) Answer() string {
	return q.answer
}

//Question : Question inteface method for Question
func (q ChoiceQuestion) Question() string {
	result := q.question + "\n"
	for i, c := range q.choices {
		result += fmt.Sprintf("%v) %v \n", i+1, c)
	}
	result += "Enter the number of the correct answer"
	return result
}

//NewSimpleQuestion : create new  question
func NewChoiceQuestion(q string, a string, ch ...string) *ChoiceQuestion {
	return &ChoiceQuestion{
		question: q,
		answer:   a,
		choices:  ch,
	}
}
