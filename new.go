package main

import (
	"fmt"
	"github.com/nanavortex/nob/blogmngr"
)

var cmdNew = &Command{
	Run: runNew,
	Usage: "new [<entry title>]",
	Desc: "Create a new blog post/entry as a markdown draft",
}

func init() {
	NobRunner.AppendCmd("new", cmdNew)
}

func runNew(params *[]string) int {
	if !blogmngr.IsBlog() {
		fmt.Println("The current directory does not seem to be a nob blog.")
		return 1
	}

	title := "New Post"
	if len(*params) > 0 {
		title = (*params)[0]
	}

	fmt.Printf("Creating draft for '%s'\n", title)
	if !blogmngr.CreateDraftAndEdit(title) {
		fmt.Println("Please choose another title (that entry already exists)")
		return 1
	}
	fmt.Println("Remember to run 'nob build' to apply changes to your blog.")
	return 0
}

