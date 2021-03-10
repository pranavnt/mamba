package mamba

import (
	"fmt"
	"testing"
)

func test(t *testing.T) {
	app := New()
	app.AddCommand("run {fileName}", runFile)
	app.AddCommand("deploy {directory}", deploy)
	app.AddCommand("help {command}", help)
	app.Run([]string{"ace", "run", "helloWorld.py"})
	app.Run([]string{"ace", "deploy", "src"})
	app.Run([]string{"ace", "help", "run"})
}

func runFile(params Dict) {
	fmt.Println("About to run " + params["fileName"].(string))
	// other code for running the code
}

func deploy(params Dict) {
	fmt.Println("About to deploy " + params["directory"].(string))
	// other code for deployment
}

func help(params Dict) {
	fmt.Println("Documentation on " + params["command"].(string))
	// show documentation
}
