package main

import (
	"mvc/routes"
)

func main() {
	router := routes.GetRouter()
	router.Run(":3000")
}
