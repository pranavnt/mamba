<h1 align="center">Mamba</h1> 
<p align="center">
    <img src="./.github/mamba.png" width="15%" />
</p>
<h2 align="center">Build command line apps in Go with ease</h2>

## Installation 💻

To install mamba, run the following command in terminal:

```bash
go get github.com/pranavnt/mamba
```

To use it, include it in your imports:

```go
import (
    "github.com/pranavnt/mamba"
)
```

## Usage 📝

For the usage example, we'll create a command line app called `ace`, which is used for deploying apps.

First, we will want to initialize a CLI app, which is done through:

```go
app := mamba.New()
```

Next, we will want to add commands to our app, which you do through `AddCommand`:

```go
app.AddCommand("deploy {fileName}", deploy)
```

The command above creates the deploy command, which takes in a parameter for `fileName` (the user will enter this). For example, this would handle `ace run hello.py`.

Now, let's write code for the function that handles this command, called `deploy`. This function must accept one parameter (of type `mamba.Dict`).

```go
func deploy(params mamba.Dict) {
	fmt.Println("About to deploy " + params["fileName"].(string))
	// code for deployment
}
```

Here is a sample function where you can read and print the fileName. You can later use this parameter to actually deploy the project in this app.

Finally, in order to run the code, you need to add:

```go
app.Run(os.Args)
```

Here, you are actually running the app, and `os.Args` are the command line arguments that are parsed by mamba. You're done!

The full code:

```go
package main

import (
	"fmt"
    "os"
	"github.com/pranavnt/mamba"
)

func main() {
	app := mamba.New()
	app.AddCommand("deploy {directory}", deploy)
	app.Run(os.Args)
}

func deploy(params mamba.Dict) {
	fmt.Println("About to deploy " + params["directory"].(string))
	// other code for deployment
}
```

## Design Goals 🎨

### Simplicity

### Readability
