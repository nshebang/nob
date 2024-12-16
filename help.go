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

Wiki: https://github.com/nshebang/nob/wiki
`

func PrintHelp() {
	fmt.Print(helpMsg)	
}

