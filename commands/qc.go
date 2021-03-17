package main

import (
	"flag"
	"fmt"
)

var commands []command

func main() {
	var command string
	var help bool
	var about bool

	flag.BoolVar(&help, "h", false, "Shows list of all the available commands.")
	flag.BoolVar(&help, "help", false, "Shows list of all the available commands.")
	flag.StringVar(&command, "c", "", "Shows information of a command.")
	flag.StringVar(&command, "command", "", "Shows information of a command.")
	flag.BoolVar(&about, "a", false, "Shows about information.")
	flag.BoolVar(&about, "about", false, "Shows about information.")
	flag.Parse()

	if help {
		helpBasic()
	} else if command != "" {
		helpCommand(command)
	} else if about {
		fmt.Println("\nQuick Commands is a collection of commands to use in your terminal/command prompt.\nFeel free to add custom commands to the \033]8;;https://github.com/Yadiiiig/quick_commands\033\\repository\033]8;;\033\\.\nCreated by \033]8;;https://github.com/Yadiiiig\033\\Yadiiiig\033]8;;\033\\")
	}
}

func helpBasic() {
	for i := range commands {
		fmt.Printf("%s - %s\n", commands[i].Name, commands[i].Information)
	}
}

func helpCommand(command string) {
	for i := range commands {
		if commands[i].Name == command {
			fmt.Printf("\nCommand: %s\nParams: %s\nInformation: %s\nExample: %s\n", commands[i].Name, commands[i].Params, commands[i].Information, commands[i].Example)
		}
	}
}

type command struct {
	Name        string
	Params      string
	Example     string
	Information string
}

func init() {
	commands = []command{
		command{
			Name:        "searchfile",
			Params:      "file-/folder name",
			Example:     "searchfile foo.bar",
			Information: "Finds all the corresponding paths of files and folder with that specific name.",
		},
		command{
			Name:        "timer",
			Params:      "minutes",
			Example:     "timer 120",
			Information: "Starts a countdown timer for specified time.",
		},
		command{
			Name:        "so",
			Params:      "-l/-limit int | search term",
			Example:     "so -l 5 Exit vim | s Exit vim",
			Information: "Stackoverflow search, limit flag is optional. From default it returns 7 links.",
		},
		command{
			Name:        "lock",
			Params:      "",
			Example:     "lock",
			Information: "Locks the PC, windows only.",
		},
		command{
			Name:        "qc",
			Params:      "-h/-help | -c/-command name | -a/-about",
			Example:     "qc -h | qc -c searchile | qc -about",
			Information: "Shows all the commands, or get specific information about a command. About shows some information about Quick Commands.",
		},
	}
}
