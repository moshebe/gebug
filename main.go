package main

import "gebug/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
