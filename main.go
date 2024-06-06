package main

import (
	"github.com/Vaniog/golang-gh-actions/internal/app"
	"log"
)

func main() {
	a := app.New()
	if err := a.Run(); err != nil {
		log.Panicln(err)
	}
}
