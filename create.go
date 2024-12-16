package main

import (
	"fmt"
	"github.com/nshebang/nob/blogmngr"
)

var cmdCreate = &Command{
	Run: runCreate,
	Usage: "create <blog_directory_name> [--url <site_url> | -u <site_url>]",
	Desc: "Create a new blog",
}

func init() {
	NobRunner.AppendCmd("create", cmdCreate)
}

var gitignoreText = `---
NEOCITIES CLI USERS:
It is recommended to add the .nob/ directory of your blog to a
.gitignore file in your blog root directory. This will prevent
Neocities CLI from uploading Markdown drafts and HTML templates.

Run the following command to add .nob/ to .gitignore:

echo "%s/.nob/" >> .gitignore
---
`

func runCreate(params *[]string) int {
	if len(*params) == 0 {
		fmt.Println("Error: Missing blog directory name")
		fmt.Println("Usage: nob create <dirname> [-u <site_url>]")
		return 1
	}

	dirname := (*params)[0]
	siteUrl := fmt.Sprintf("https://my-website.neocities.org/%s", dirname)
	fmt.Printf("Creating blog in '%s'\n", dirname)
	
	if len(*params) == 3 {
		if (*params)[1] != "--url" && (*params)[1] != "-u" {
			fmt.Printf("Invalid flag %s\n", (*params)[1])
			return 1
		}

		siteUrl = (*params)[2]
		fmt.Printf("Blog URL is %s\n", siteUrl)
	}

	success := blogmngr.CreateBlog(dirname, siteUrl)
	if !success {
		fmt.Printf("The directory '%s' already exists\n", dirname)
		return 1
	}

	fmt.Printf(gitignoreText, dirname)
	fmt.Println("Blog successfully created. Remember to customize your templates!")
	return 0
}

