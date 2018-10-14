package task

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type Task struct {
	Id  string `json:"id"`
	Cmd string `json:"cmd"`
}

type client struct {
	whitelistedCmds []string
}

type Client interface {
	Exec(msg []byte) error
}

func New(whitelistedCmds []string) Client {
	return &client{
		whitelistedCmds: whitelistedCmds,
	}
}

func (c *client) Exec(msg []byte) error {
	fmt.Println(msg)
	t, err := Parse(msg)
	if err != nil {
		return errors.Wrap(err, "could not parse string to task")
	}

	fmt.Println(t.Cmd)

	return nil
}

func Parse(b []byte) (Task, error) {
	var t Task
	err := json.Unmarshal(b, &t)
	if err != nil {
		return Task{}, errors.Wrap(err, "could not unmarshal json")
	}

	return t, nil
}

func String(t Task) string {
	s, err := json.Marshal(t)
	if err != nil {
		fmt.Println("could not marshal json")
	}
	return string(s)
}

// Running tasks confirm white listed commands
