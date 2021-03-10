package main

import (
	"fmt"
)

func New() {

}

func addCommand(cmd string, fn function) {
	fmt.Println(cmd)
}

func Run() {

}

type dict map[string]interface{}

type function func(args dict)
