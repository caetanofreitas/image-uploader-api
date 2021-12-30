package main

import (
	"uploader/cmd"
	env "uploader/environment"
)

func main() {
	env.ValidateEnvironment()
	cmd.ExecuteRouter()
}
