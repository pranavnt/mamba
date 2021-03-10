package main

import (
	"fmt"
)

func main() {
	app := New()
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

func (cli cli) addCommand(cmd string, fn function) {
	fmt.Println(cmd)
}

func Run() {

}

type cli struct {
	commands []cmd
}

type cmd struct {
	cmd    string
	length int
	fn     function
}

type Dict map[string]interface{}

type function func(args Dict)
