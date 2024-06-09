package main

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io"
	"log"
	"os"

	"github.com/ryanbabida/tcpcache/cache"
	"github.com/ryanbabida/tcpcache/server"
)

func main() {
	hashFunc := func(key string) int {
		h := fnv.New32a()
		h.Write([]byte(key))
		return int(h.Sum32())
	}

	c, err := cache.NewCache[string, any](5, hashFunc)
	if err != nil {
		log.Fatalln("failed to initialize cache")
	}

	cfg, err := readJSONFile("settings.json")
	if err != nil {
		log.Println("unable to read JSON config file, will use default config instead: %w", err)
	}

	s := server.NewServer(cfg, c)
	s.Run()
}

func readJSONFile(filename string) (*server.Config, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to open filename '%s': %w", filename, err)
	}

	b, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %w", err)
	}

	var c server.Config
	err = json.Unmarshal(b, &c)
	if err != nil {
		return nil, fmt.Errorf("unable to parse json file: %w", err)
	}

	return &c, nil
}
