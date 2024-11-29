package main

import (
	"log"
	"os"

	"app.com/cli/app"
)

func main() {
	app := app.CreateApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
