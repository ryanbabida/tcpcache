package server

import "testing"

func TestConfig_NewConfig(t *testing.T) {
	cfg := NewConfig()

	if *cfg.Port != defaultPort {
		t.Errorf("expected: %v, actual: %v", defaultPort, *cfg.Port)
	}
}

func TestConfig_WithPortOpts(t *testing.T) {
	cfg := NewConfig(WithPort("1234"))

	if *cfg.Port != "1234" {
		t.Errorf("expected: %v, actual: %v", "1234", *cfg.Port)
	}
}
