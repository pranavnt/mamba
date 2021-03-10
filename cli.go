package main

import (
	"fmt"
	"strings"
)

func main() {
	app := New()
	app.addCommand("asteroid run hi", hi)
	app.addCommand("asteroid run hi", hi)
	app.addCommand("asteroid run hi", hi)
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
	fmt.Println(cli.commands)
}

func Run() {

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
