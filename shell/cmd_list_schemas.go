package shell

import (
	"io"

	"github.com/swanwish/go-common/logs"
)

func init() {
	Commands = append(Commands, &commandListSchema{})
}

type commandListSchema struct{}

func (c *commandListSchema) Name() string {
	return ".schema"
}

func (c *commandListSchema) Help() string {
	return `list schema of the tables on the database

	.schema
	.schema pattern

`
}

func (c *commandListSchema) Do(args []string, commandEnv *CommandEnv, writer io.Writer) (err error) {
	querySql := "select sql from sqlite_schema where type='table'"
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
	options := printOptions{showHeader: false, lineSeparator: "\n"}
	queryResult.ShowResult(options)
	return nil
}
