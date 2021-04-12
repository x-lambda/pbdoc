package main

import (
	"pbdoc/router"
)

func main() {
	g := router.GetRouter()

	g.Run()
}
