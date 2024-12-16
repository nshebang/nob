package main

import (
	"fmt"
	"time"
	"github.com/nshebang/nob/blogmngr"
)

var cmdBuild = &Command{
	Run: runBuild,
	Usage: "build [--autodelete | -d]",
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

	autodelete := false
	if len(*params) > 0 {
		if (*params)[0] == "--autodelete" || (*params)[0] == "-d" {
			autodelete = true
		}
	}

	fmt.Println("Building public HTML files + RSS feed")
	if autodelete {
		fmt.Println("Entries without a Markdown file will be deleted!")
	}
	start := time.Now().UnixNano() / int64(time.Millisecond)

	blogmngr.BuildStatic(autodelete)

	end := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Printf("Done! (%d ms)\n", end - start)

	return 0
}

