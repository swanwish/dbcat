package shell

import (
	"io"

	"github.com/jmoiron/sqlx"
	"github.com/swanwish/dbcat/common"
	"github.com/swanwish/go-common/logs"
)

type ShellOptions struct {
}

type CommandEnv struct {
	env    map[string]string
	option *ShellOptions
	db     *sqlx.DB
}

type command interface {
	Name() string
	Help() string
	Do([]string, *CommandEnv, io.Writer) error
}

var (
	Commands = []command{}
)

func NewCommandEnv(db *sqlx.DB) *CommandEnv {
	return &CommandEnv{db: db}
}

func (ce *CommandEnv) executeSql(querySql string, params []interface{}) (QueryResult, error) {
	if ce == nil {
		logs.Errorf("The command env is nil")
		return QueryResult{}, common.ErrInvalidParameter
	}

	if querySql == "" {
		logs.Errorf("The query sql is empty")
		return QueryResult{}, common.ErrInvalidParameter
	}

	rows, err := ce.db.Queryx(querySql, params...)
	if err != nil {
		logs.Errorf("Failed to query sql, the error is %#v", err)
		return QueryResult{}, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		logs.Errorf("Failed to get columns, the error is %#v", err)
		return QueryResult{}, err
	}
	colNumber := len(columns)
	queryResult := QueryResult{Columns: columns}
	for rows.Next() {
		values := make([]interface{}, colNumber)
		for i := range values {
			values[i] = new(interface{})
		}
		err = rows.Scan(values...)
		if err != nil {
			logs.Errorf("Failed to scan, the error is %#v", err)
			return queryResult, err
		}
		queryResult.Rows = append(queryResult.Rows, values)
		// for index, value := range values {
		// 	logs.Debugf("%s: %v", columns[index], *(value.(*interface{})))
		// }
		// fmt.Println()
	}
	return queryResult, nil
}
