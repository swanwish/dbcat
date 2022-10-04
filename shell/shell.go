package shell

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/peterh/liner"
	"github.com/swanwish/dbcat/common"
	"github.com/swanwish/go-common/logs"
)

var (
	line            *liner.State
	historyFileName = fmt.Sprintf(".%s", common.AppName)
	historyPath     string
	commands        = make([]string, 0)
	mode            = "column"
)

func RunShell(dbPath, logPath string) error {
	historyPath = path.Join(logPath, historyFileName)
	db, err := sqlx.Connect("sqlite3", dbPath)
	if err != nil {
		logs.Errorf("Failed to open sqlite db from path %s, the error is %#v", dbPath, err)
		return err
	}
	commandEnv := NewCommandEnv(db)

	line = liner.NewLiner()
	defer prepareExit()

	line.SetCtrlCAborts(true)
	line.SetTabCompletionStyle(liner.TabPrints)

	setCompletionHandler()
	loadHistory()

	cmdLines := make([]string, 0)
	isMultiLines := false
	for {
		prompt := fmt.Sprintf("%s > ", common.AppName)
		if isMultiLines {
			prompt = fmt.Sprintf("%*s > ", len(common.AppName), "...")
		}
		cmd, err := line.Prompt(prompt)
		if err != nil {
			if err != io.EOF {
				fmt.Printf("%v\n", err)
			}
			return err
		}

		cmd = strings.TrimSpace(cmd)
		if cmd == "" {
			continue
		}

		if cmd == "\\c" {
			cmdLines = cmdLines[:0]
			isMultiLines = false
			continue
		}

		if !strings.HasPrefix(cmd, ".") && !strings.HasSuffix(cmd, ";") {
			cmdLines = append(cmdLines, cmd)
			isMultiLines = true
			continue
		}

		if isMultiLines {
			cmdLines = append(cmdLines, cmd)
			cmd = strings.Join(cmdLines, " ")
			cmdLines = cmdLines[:0]
			isMultiLines = false
		}

		if err = processCmd(cmd, commandEnv); errors.Is(err, common.ErrExit) {
			break
		}
	}

	return nil
}

func processCmd(cmd string, commandEnv *CommandEnv) error {
	line.AppendHistory(cmd)

	args := strings.Split(cmd, " ")
	if len(args) == 0 {
		return common.ErrInvalidParameter
	}

	cmd = args[0]
	if cmd == ".help" || cmd == ".?" {
		printHelp(args)
		return nil
	} else if cmd == ".exit" || cmd == ".quit" || cmd == ".q" {
		return common.ErrExit
	}

	if !strings.HasPrefix(cmd, ".") {
		cmd = ".query"
	}

	for _, command := range Commands {
		if command.Name() == cmd {
			return command.Do(args, commandEnv, os.Stdout)
		}
	}

	fmt.Printf("unknown command: %s", cmd)
	return common.ErrInvalidParameter
}

func printGenericHelp() {
	msg :=
		`Type:	"help <command>" for help on <command>. Most commands support "<command> -h" also for options. 
`
	fmt.Print(msg)

	for _, c := range Commands {
		helpTexts := strings.SplitN(c.Help(), "\n", 2)
		fmt.Printf("  %-30s\t# %s \n", c.Name(), helpTexts[0])
	}
}

func printHelp(cmds []string) {
	args := cmds[1:]
	if len(args) == 0 {
		printGenericHelp()
	} else if len(args) > 1 {
		fmt.Println()
	} else {
		cmd := strings.ToLower(args[0])

		for _, c := range Commands {
			if c.Name() == cmd {
				fmt.Printf("  %s\t# %s\n", c.Name(), c.Help())
			}
		}
	}
}

func setCompletionHandler() {
	line.SetCompleter(func(line string) (c []string) {
		for _, i := range Commands {
			if strings.HasPrefix(i.Name(), strings.ToLower(line)) {
				c = append(c, i.Name())
			}
		}
		return
	})
}

func loadHistory() {
	if f, err := os.Open(historyPath); err == nil {
		line.ReadHistory(f)
		f.Close()
	}
}

func prepareExit() {
	saveHistory()
	line.Close()
}

func saveHistory() {
	if f, err := os.Create(historyPath); err != nil {
		fmt.Printf("Error creating history file: %v\n", err)
	} else {
		for _, cmd := range commands {
			line.AppendHistory(cmd)
		}
		if _, err = line.WriteHistory(f); err != nil {
			fmt.Printf("Error writing history file: %v\n", err)
		}
		f.Close()
		logs.Debugf("History saved to %s", historyPath)
	}
}
