package todo

type TodoEntity struct {
	Title string `json:"title"`
}

type TodoError struct {
	Msg string `json:"msg"`
}
