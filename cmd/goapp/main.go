package main

import (
	"github.com/randallmlough/goapp/command"
)

func main() {
	if err := command.Execute(); err != nil {
		panic(err)
	}
}
