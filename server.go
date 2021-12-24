package main

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello World."
	})

	m.Post("/call/:title", func(params martini.Params) string {
		return "Saved new call"
	})

	m.Run()

}
