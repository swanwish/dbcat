package main

import (
	"os"
	"runtime"

	"github.com/swanwish/dbcat/cmd"
	"github.com/swanwish/go-common/logs"
	"github.com/urfave/cli/v2"
)

const AppVersion = "0.0.1"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := cli.NewApp()
	app.Name = "dbhelper"
	app.Usage = "The dbhelper service"
	app.Version = AppVersion
	app.Commands = []*cli.Command{
		&cmd.ShellCmd,
	}
	app.Action = cmd.ShellCmd.Action
	app.Flags = append(app.Flags, cmd.ShellCmd.Flags...)
	if err := app.Run(os.Args); err != nil {
		logs.Errorf("Failed to run application, the error is %#v", err)
	}
}
