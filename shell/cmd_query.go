package shell

import (
	"io"
	"strings"

	"github.com/swanwish/dbcat/common"
	"github.com/swanwish/go-common/logs"
)

func init() {
	Commands = append(Commands, &commandQuery{})
}

type commandQuery struct{}

func (c *commandQuery) Name() string {
	return ".query"
}

func (c *commandQuery) Help() string {
	return `Execute query sql and print the result, .query can be ignored

	.query <sql>
	<sql>

`
}

func (c *commandQuery) Do(args []string, commandEnv *CommandEnv, writer io.Writer) (err error) {
	if len(args) == 0 {
		err = common.ErrInvalidParameter
		return
	}

	if args[0] == c.Name() {
		args = args[1:]
	}
	querySql := strings.Join(args, " ")
	params := make([]interface{}, 0)
	queryResult, err := commandEnv.executeSql(querySql, params)
	if err != nil {
		logs.Errorf("Failed to execute querySql %s, the error is %#v", querySql, err)
		return err
	}
	queryResult.ShowResult(NewPrintOptions())
	return nil
}
