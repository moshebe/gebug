package main

import (
	"github.com/moshebe/gebug/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
