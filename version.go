package main

import (
	"fmt"
)

var cmdVersion = &Command{
	Run: runVersion,
	Usage: "version",
	Desc: "Show nob version and exit",
}

const version = "1.1.0"

func init() {
	NobRunner.AppendCmd("version", cmdVersion)
}

func runVersion(params *[]string) int {
	fmt.Printf("Version %s\n", version)
	return 0
}

