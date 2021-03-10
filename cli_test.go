package main

import (
	"fmt"
	"os"
	"testing"
)

func test(t *testing.T) {
	app := New()
	app.addCommand("{testing}", hi)
	app.addCommand("hi {d}", hi)
	app.addCommand("run hi 2", hi)
	app.Run(os.Args)
	fmt.Println()
}

func hi(hi Dict) {
	fmt.Println("hi")
	fmt.Println("hi")
}
