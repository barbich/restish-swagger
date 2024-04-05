package main

import (
	"github.com/barbich/restish-swagger/swagger"
	"os"

	"github.com/barbich/restish/cli"
	// "github.com/google/martian/log"
)

var version string = "dev"
var commit string
var date string

func main() {
	if version == "dev" {
		// Try to add the executable modification time to the dev version.
		filename, _ := os.Executable()
		if info, err := os.Stat(filename); err == nil {
			version += "-" + info.ModTime().Format("2006-01-02-15:04")
		}
	}

	cli.Init("restish-swagger", version)

	// Register default encodings, content type handlers, and link parsers.
	cli.Defaults()

	// Register format loaders to auto-discover API descriptions
	// cli.AddLoader(openapi.New())
	cli.AddLoader(swagger.New())

	// Run the CLI, parsing arguments, making requests, and printing responses.
	if err := cli.Run(); err != nil {
		os.Exit(1)
	}

	// Exit based on the status code of the last request.
	os.Exit(cli.GetExitCode())
}
