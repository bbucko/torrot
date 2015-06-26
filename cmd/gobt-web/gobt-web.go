package main

import (
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	m.Get("/:name", func() string {
		return "Hello World + 1"
	})
	m.Run()
}
