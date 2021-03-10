package mamba

import (
	"fmt"
	"strings"
)

func main() {}

func New() cli {
	var cmds []cmd
	return cli{
		commands: cmds,
	}
}

func (cli *cli) AddCommand(command string, fn function) {
	cli.commands = append(cli.commands, cmd{
		cmd:    command,
		arr:    strings.Split(command, " "),
		length: len(strings.Split(command, " ")),
		fun:    fn,
	})
}

func (cli *cli) Run(args []string) {
	//creating dict
	param := make(Dict)

	// going through all commands
	for _, command := range cli.commands {
		// checking if lengths match up
		if command.length == len(args)-1 {
			// going through words in command
			for i := 0; i < command.length; i++ {
				// checking if word is like {this}
				if checkArg(command.arr[i]) {
					// fmt.Println(command.arr[i][1 : len(command.arr[i])-1])
					param[command.arr[i][1:len(command.arr[i])-1]] = args[i+1]
				} else {
					// if the commands are not the same move to the next one
					if command.arr[i] != args[i+1] {
						break
					}
				}

				// running command!
				if i+1 == command.length {
					command.fun(param)
					return
				}
			}
		}
	}
	fmt.Println("Command not found :(")
}

func checkArg(arg string) bool {
	if arg[0:1] == "{" && arg[len(arg)-1:] == "}" {
		return true
	} else {
		return false
	}
}

type cli struct {
	commands []cmd
}

type cmd struct {
	cmd    string
	arr    []string
	length int
	fun    function
}

type Dict map[string]interface{}

type function func(args Dict)
