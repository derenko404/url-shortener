package main

import (
	"fmt"

	"github.com/url-shortener/internal/cache"
	"github.com/url-shortener/web/server"
)

func main() {
	cache := cache.New()
	server := server.New(":8080", cache)
	err := server.Start()

	if err != nil {
		panic(fmt.Sprintf("cannot start server %v", err))
	}
}
