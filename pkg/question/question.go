package question

type Question struct {
	Content string
	Answer  int
}

var Questions []Question

// AddQuestion adds a question to the questions slice
func Add(c string, a int) []Question {
	question := Question{
		Content: c,
		Answer:  a,
	}

	Questions = append(Questions, question)

	return Questions
}
