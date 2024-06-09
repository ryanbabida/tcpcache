package main

import (
	"hash/fnv"
	"log"

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

	cfg, err := server.ReadJSONFile("settings.json")
	if err != nil {
		log.Println("unable to read JSON config file, will use default config instead: %w", err)
	}

	s := server.NewServer(cfg, c)
	s.Run()
}
