package main

import (
	"fmt"
)

type Command struct {
	Run	func(params *[]string) int
	Usage	string
	Desc	string
}

func (c *Command) PrintUsageAndDesc() {
	fmt.Printf("Usage: nob %s\n\n", c.Usage)
	fmt.Printf("%s\n", c.Desc)
}

