package main

import (
	"fmt"
)

var cmdHelp = &Command{
	Run: runHelp,
	Usage: "help [<command>] | [--more]",
	Desc: "Show help about nob or about a specific command",
}

func init() {
	NobRunner.AppendCmd("help", cmdHelp)
}

func runHelp(params *[]string) int {
	if len(*params) == 0 {
		PrintHelp()
		return 0
	}
	if (*params)[0] == "--more" {
		printMoreHelp()
		return 0
	}
	cmdname := (*params)[0]
	cmd := NobRunner.commands[cmdname]
	if cmd != nil {
		cmd.PrintUsageAndDesc()	
	} else {
		fmt.Println("Command not found.")
		return 1
	}
	return 0
}

var helpMsg = `nob is a neocities-oriented blog manager.
Usage: nob <command> [<params>]

Commands:
 create		Create a new blog. You must type the dirname
 new		Create a new entry for your blog
 build		Apply changes to your blog HTML files
 help		Show help about a command or about nob
 version	Show version and exit

Use nob help --more to get more help.
`

func PrintHelp() {
	fmt.Print(helpMsg)	
}

var moreHelp = `nob user help
---
* What is nob?
 nob is a blog manager (and static site generator) that's truly easy to use.
 It is intended to be used by Neocities users, but it works even if you're
 using GitHub pages or really any static website hosting service.

* How to edit an entry
 Open the entry file in the nob/drafts/ directory and edit it using your
 preferred text editor.

* How to delete an entry
 Just delete it from nob/drafts/ and delete its HTML file, too. Then, run
 nob build.

* How to apply changes
 Run nob build

* How to publish your blog on Neocities
 Just upload the files in your blog directory, except for all in nob/, since
 those files are only used to easily edit your entries and HTML if you need
 to. If you are using the Neocities CLI program, run: neocities push.

* My blog looks ugly! How can I make it look pretty?
 Make a CSS stylesheet and add it to nob/templates/*.html using a <link> tag.
 For example: <link rel="stylesheet" href="my_theme.css" />
 Just remember that the <link> tag must be placed inside the <head> tag.
 And remember to run 'nob build' to apply changes, too.
`

func printMoreHelp() {
	fmt.Print(moreHelp)
}

