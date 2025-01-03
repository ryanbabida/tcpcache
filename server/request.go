package server

import "fmt"

type request[K comparable, V any] struct {
	Action Action `json:"action"`
	Key    *K     `json:"key"`
	Value  *V     `json:"value"`
}

func (r *request[K, V]) isValid() error {
	if err := r.Action.IsValid(); err != nil {
		return fmt.Errorf("invalid request with bad action: %v", err)
	}

	if r.Key == nil {
		return fmt.Errorf("key was not provided")
	}

	if r.Action == Get && r.Value != nil {
		return fmt.Errorf("GET action cannot have a value defined")
	}

	if r.Action == Set && r.Value == nil {
		return fmt.Errorf("SET action must have a value provided")
	}

	return nil
}
