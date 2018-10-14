package task

type Task struct {
	Id  string `json:"id"`
	Cmd string `json:"cmd"`
}

// Running tasks confirm white listed commands
