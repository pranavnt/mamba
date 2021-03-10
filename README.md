<h1 align="center">Mamba</h1> 
<p align="center">
    <img src="./.github/mamba.png" width="15%" />
</p>
<h2 align="center">Build command line apps in Go with ease</h2>

## Installation

To install mamba, run the following command in terminal:

```bash
go get github.com/pranavnt/mamba
```

To use it, include it in your imports:

```go
import (
    github.com/pranavnt/mamba
)
```

## Usage

For the usage example, we'll create a command line app called `ace`, which is used for deploying apps.

First, we will want to initialize a CLI app, which is done through:

```go
app := New()
```

Next, we will want to add commands to our app, which you do through `addCommand`:

```go
app.addCommand("deploy {fileName}", deploy)
```

The command above creates the deploy command, which takes in a parameter for `fileName` (the user will enter this). For example, this would handle `ace run hello.py`.

Now, let's write code for the function that handles this command, called `deploy`. This function must accept one parameter (of type `mamba.Dict`).

```go
func deploy(params mamba.Dict) {
	fmt.Println("About to deploy " + params["fileName"].(string))
	// code for deployment
}
```

Here, this function reads the fileName parameter

## Design Goals
