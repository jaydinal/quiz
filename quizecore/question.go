package quizecore

type Question interface {
	Question() string
	Validate(string) bool
	Answer() string
}
