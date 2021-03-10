package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	app := New()
	app.addCommand("{testing}", hi)
	app.addCommand("hi {d}", hi)
	app.addCommand("run hi 2", hi)
	app.Run(os.Args)
}

func hi(hi Dict) {
	fmt.Println("hi")
}

func New() cli {
	var cmds []cmd
	return cli{
		commands: cmds,
	}
}

func (cli *cli) addCommand(command string, fn function) {
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
			for i := 0; i <= command.length-1; i++ {
				// checking if word is like {this}
				if checkArg(command.arr[i]) {
					// fmt.Println(command.arr[i][1 : len(command.arr[i])-1])
					param[command.arr[i][1:len(command.arr[i])-1]] = args[i+1]
					fmt.Println(param)
				} else {
					if command.arr[i] == args[i+1] {
						fmt.Println("continue")
						continue
					} else {
						fmt.Println("break")
						break
					}
				}
			}
		}
	}
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
