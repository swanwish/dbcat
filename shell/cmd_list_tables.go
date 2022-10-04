package shell

import (
	"io"

	"github.com/swanwish/go-common/logs"
)

func init() {
	Commands = append(Commands, &commandListTable{})
}

type commandListTable struct{}

func (c *commandListTable) Name() string {
	return ".tables"
}

func (c *commandListTable) Help() string {
	return `list tables on the database

	.tables
	.tables pattern

`
}

func (c *commandListTable) Do(args []string, commandEnv *CommandEnv, writer io.Writer) (err error) {
	querySql := "select name from sqlite_schema where type='table'"
	params := make([]interface{}, 0)
	if len(args) > 1 {
		querySql += " and name like ?"
		params = append(params, args[1])
	}
	queryResult, err := commandEnv.executeSql(querySql, params)
	if err != nil {
		logs.Errorf("Failed to execute sql, the error is %#v", err)
		return err
	}
	queryResult.ShowResult(NewPrintOptions())
	return nil
}
