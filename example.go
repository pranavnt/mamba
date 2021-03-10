package main

import (
	"fmt"
	"os"
)

func main() {
	app := New()
	app.addCommand("run {fileName}", runFile)
	app.addCommand("deploy {directory}", deploy)
	app.addCommand("help {command}", help)
	app.Run(os.Args)
}

func runFile(params Dict) {
	fmt.Println("About to run " + params["fileName"].(string))
	// other code
}

func deploy(params Dict) {
	fmt.Println("About to deploy " + params["directory"].(string))
	// other code
}

func help(params Dict) {
	fmt.Println("Documentation on " + params["command"].(string))
	// show documentation
}
