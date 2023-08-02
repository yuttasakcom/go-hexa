package todo

type Todo struct {
	Title string `json:"title"`
}

type TodoError struct {
	Msg string `json:"msg"`
}
