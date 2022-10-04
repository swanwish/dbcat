package cmd

import "github.com/urfave/cli/v2"

func stringFlag(name, value, usage string) cli.Flag {
	return &cli.StringFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func boolFlag(name, usage string) cli.Flag {
	return &cli.BoolFlag{
		Name:  name,
		Usage: usage,
	}
}

func intFlag(name string, value int, usage string) cli.Flag {
	return &cli.IntFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}
