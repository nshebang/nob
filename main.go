package main

import (
	"os"
)

var NobRunner = NewRunner()

func main() {
	err := NobRunner.Execute()
	os.Exit(err)
}

