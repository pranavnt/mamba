package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
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
		length: len(strings.Split(command, " ")),
		fun:    fn,
	})
}

func (cli *cli) Run(args []string) {
	for _, command := range cli.commands {
		if command.length == len(os.Args)-1 {
			fmt.Println(command.cmd)
		}
	}
}

type cli struct {
	commands []cmd
}

type cmd struct {
	cmd    string
	length int
	fun    function
}

type Dict map[string]interface{}

type function func(args Dict)
