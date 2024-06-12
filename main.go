package main

import (
	"flag"
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

	opts := []func(*server.Config){}

	port := flag.String("port", "", "listening address for server")
	flag.Parse()

	if len(*port) > 0 {
		opts = append(opts, server.WithPort(*port))
	}

	cfg := server.NewConfig(opts...)
	s := server.NewServer(c, *cfg)
	s.Run()
}
