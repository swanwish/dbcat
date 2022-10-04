package shell

import (
	"fmt"
	"io"

	"github.com/swanwish/dbcat/common"
	"github.com/swanwish/go-common/logs"
)

func init() {
	Commands = append(Commands, &commandListColumns{})
}

type commandListColumns struct{}

func (c *commandListColumns) Name() string {
	return ".columns"
}

func (c *commandListColumns) Help() string {
	return `list columns of the table on the database

	.columns table

`
}

func (c *commandListColumns) Do(args []string, commandEnv *CommandEnv, writer io.Writer) (err error) {
	if len(args) < 2 {
		logs.Errorf("The table name does not specified")
		return common.ErrInvalidParameter
	}
	querySql := fmt.Sprintf("pragma table_info(%s)", args[1])
	params := make([]interface{}, 0)
	queryResult, err := commandEnv.executeSql(querySql, params)
	if err != nil {
		logs.Errorf("Failed to execute sql, the error is %#v", err)
		return err
	}
	queryResult.ShowResult(NewPrintOptions())
	return nil
}
