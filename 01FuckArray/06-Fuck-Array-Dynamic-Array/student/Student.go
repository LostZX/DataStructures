package student

type Student struct {
	name  string
	score int
}

func Instance(name string, score int) *Student {
	return &Student{name, score}
}
