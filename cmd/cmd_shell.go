package cmd

import (
	"os"

	"github.com/swanwish/dbcat/common"
	"github.com/swanwish/dbcat/shell"
	"github.com/swanwish/go-common/logs"
	"github.com/urfave/cli/v2"
)

var (
	ShellCmd = cli.Command{
		Name:        "shell",
		Usage:       "Start the shell command",
		Description: "This command will start the shell command",
		Action:      shellAction,
		Flags: []cli.Flag{
			stringFlag("dbPath", "", "The path of the db file"),
			stringFlag("logPath", os.TempDir(), "The path to store history file"),
		},
	}
)

func shellAction(c *cli.Context) error {
	dbPath := c.String("dbPath")
	if dbPath == "" {
		args := c.Args()
		if args.Len() == 1 {
			dbPath = args.First()
		} else {
			logs.Errorf("The dbPath does not specified")
			return common.ErrInvalidParameter
		}
	}
	logPath := c.String("logPath")
	return shell.RunShell(dbPath, logPath)
}
