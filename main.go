package main

import (
	"uploader/cmd"
	env "uploader/environment"
)

func init() {
	env.ValidateEnvironment()
}

func main() {
	cmd.ExecuteRouter()
}
