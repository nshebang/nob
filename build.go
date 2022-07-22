package main

import (
	"fmt"
	"time"
	"github.com/nanavortex/nob/blogmngr"
)

var cmdBuild = &Command{
	Run: runBuild,
	Usage: "build",
	Desc: "Apply changes to your blog static files",
}

func init() {
	NobRunner.AppendCmd("build", cmdBuild)
}

func runBuild(params *[]string) int {
	if !blogmngr.IsBlog() {
		fmt.Println("The current directory does not seem to be a nob blog.")
		return 1
	}

	fmt.Println("Building public HTML files + RSS feed")
	start := time.Now().UnixNano() / int64(time.Millisecond)

	blogmngr.BuildStatic()

	end := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Printf("Done! (%d ms)\n", end - start)

	return 0
}

