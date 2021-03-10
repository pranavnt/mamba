package main

import (
	"fmt"
)

func main() {
	// cli := New()
	addCommand("asteroid run hi", hi)
}

func hi(hi Dict) {
	fmt.Println("hi")
}

// func New() {
//    return 1
// }

func addCommand(cmd string, fn function) {
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
