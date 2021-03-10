package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	checkArg("{hi}")
	app := New()
	app.addCommand("hi", hi)
	app.addCommand("run hi", hi)
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
	for _, command := range cli.commands {
		if command.length == len(os.Args)-1 {
			for i := 0; i <= command.length-1; i++ {
				if checkArg(command.arr[i]) {

				} else {
					fmt.Println(args[i+1])
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
