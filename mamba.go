package mamba

import (
	"fmt"
	"strings"
)

// Creates a new CLI app with Mamba!
func New() cli {
	var cmds []cmd
	return cli{
		commands: cmds,
	}
}

// This is how you add a new command 
// The first parameter will be the specific command 
// The second command will be the function called when this command is used
func (cli *cli) AddCommand(command string, fn function) {
	cli.commands = append(cli.commands, cmd{
		cmd:    command,
		arr:    strings.Split(command, " "),
		length: len(strings.Split(command, " ")),
		fun:    fn,
	})
}

// This is how you run your CLI - the parameter should be `os.Args`, as those are the arguments to be parsed!
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

// The structure for your CLI app 
//
type cli struct {
	commands []cmd
}

type cmd struct {
	cmd    string
	arr    []string
	length int
	fun    function
}

// The type for a dictionary - This is how parameters are passed to your functions
// It's just a map[string]interface{}
type Dict map[string]interface{}

type function func(args Dict)
