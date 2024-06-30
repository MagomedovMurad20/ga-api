package main

import (
	"os"
)

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	a.Run(":8080")
}
