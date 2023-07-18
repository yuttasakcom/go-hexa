package todo

type Todo struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TodoError struct {
	Msg string `json:"msg"`
}
