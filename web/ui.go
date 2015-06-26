package web
import "github.com/go-martini/martini"

func start() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello World"
	})
	m.Run()
}