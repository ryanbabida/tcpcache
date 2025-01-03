package server

import "fmt"

type Action string

const (
	Get Action = "GET"
	Set Action = "SET"
)

func (a Action) IsValid() error {
	switch a {
	case Get, Set:
		break
	default:
		return fmt.Errorf("invalid action '%s' passed in", a)
	}

	return nil
}
