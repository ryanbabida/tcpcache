package main

import (
	"hash/fnv"
	"log"

	"github.com/ryanbabida/tcpcache/cache"
	"github.com/ryanbabida/tcpcache/server"
)

func main() {
	c, err := cache.NewCache[string, any](5, func(key string) int {
		h := fnv.New32a()
		h.Write([]byte(key))
		return int(h.Sum32())
	})

	if err != nil {
		log.Fatalln("failed to initialize cache")
	}

	s := server.NewServer(c)
	s.Run()
}
