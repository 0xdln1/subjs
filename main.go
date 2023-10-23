package main

import (
	"log"

	"github.com/0xdln1/subjs/runner/subjs"
)

func main() {
	opts := subjs.ParseOptions()
	runner := subjs.New(opts)
	err := runner.Run()
	if err != nil {
		log.Fatalf("Error running subjs: %s", err)
	}
}
