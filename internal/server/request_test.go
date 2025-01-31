package server

import "testing"

func TestRequest_ValidGet(t *testing.T) {
	action := Get
	key := "key"

	r := &request[string, int]{action, &key, nil}
	err := r.isValid()
	if err != nil {
		t.Errorf("expected: %v, actual: %v", nil, err)
	}
}

func TestRequest_ValidSet(t *testing.T) {
	action := Set
	key := "key"
	val := 1

	r := &request[string, int]{action, &key, &val}
	err := r.isValid()
	if err != nil {
		t.Errorf("expected: %v, actual: %v", nil, err)
	}
}

func TestRequest_InvalidAction(t *testing.T) {
	action := "badaction"
	key := "key"
	val := 1

	r := &request[string, int]{Action(action), &key, &val}
	err := r.isValid()
	if err == nil {
		t.Errorf("action is invalid, should have failed validation")
	}
}

func TestRequest_InvalidWithoutKey(t *testing.T) {
	action := Get

	r := &request[string, int]{Action(action), nil, nil}
	err := r.isValid()
	if err == nil {
		t.Errorf("any action is invalid without key, should have failed validation")
	}
}

func TestRequest_InvalidGetWithValue(t *testing.T) {
	action := Get
	key := "key"
	val := 1

	r := &request[string, int]{action, &key, &val}
	err := r.isValid()
	if err == nil {
		t.Errorf("get action is invalid with value, should have failed validation")
	}
}

func TestRequest_InvalidSetWithoutValue(t *testing.T) {
	action := Set
	key := "key"

	r := &request[string, int]{action, &key, nil}
	err := r.isValid()
	if err == nil {
		t.Errorf("set action is invalid without value, should have failed validation")
	}
}
