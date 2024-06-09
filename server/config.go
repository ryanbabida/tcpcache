package server

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	Port *string `json:"port"`
}

func getConfig(overrides *Config) Config {
	defaultPort := "8080"

	c := Config{
		Port: &defaultPort,
	}

	if overrides == nil {
		return c
	}

	if overrides.Port != nil {
		c.Port = overrides.Port
	}

	return c
}

func ReadJSONFile(filename string) (*Config, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to open filename '%s': %w", filename, err)
	}

	c, err := readJSONFile(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read JSON file: %w", err)
	}

	return c, nil
}

func readJSONFile(r io.Reader) (*Config, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %w", err)
	}

	var c Config
	err = json.Unmarshal(b, &c)
	if err != nil {
		return nil, fmt.Errorf("unable to parse json file: %w", err)
	}

	return &c, nil
}
