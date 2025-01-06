package main

import (
	"log"
	"os"

	"github.com/damiaoterto/ip-finder/app"
)

func main() {
	app := app.Generate()
	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
