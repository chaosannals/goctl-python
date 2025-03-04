package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/chaosannals/goctl-python/action"
	"github.com/urfave/cli/v2"
)

var (
	version  = "20250304"
	commands = []*cli.Command{
		{
			Name:   "python",
			Usage:  "generates http client for python",
			Action: action.Python,
			Flags:  []cli.Flag{},
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Usage = "a plugin of goctl to generate http client code for python."
	app.Version = fmt.Sprintf("%s %s/%s", version, runtime.GOOS, runtime.GOARCH)
	app.Commands = commands
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("goctl-python: %+v\n", err)
	}
}
