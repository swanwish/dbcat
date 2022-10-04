package shell

import (
	"fmt"
	"strings"
)

type QueryResult struct {
	Columns []string
	Rows    [][]interface{}
}

type printOptions struct {
	separator     string
	showHeader    bool
	lineSeparator string
}

func NewPrintOptions() printOptions {
	return printOptions{separator: "\t", showHeader: true, lineSeparator: ""}
}

func (queryResult QueryResult) ShowResult(options printOptions) {
	if len(queryResult.Columns) == 0 || len(queryResult.Rows) == 0 {
		return
	}

	colCount := len(queryResult.Columns)
	colLengths := make([]int, colCount)

	for index, col := range queryResult.Columns {
		colLengths[index] = len(col)
		if colLengths[index] < 10 {
			colLengths[index] = 10
		}
	}

	if options.showHeader {
		for index, col := range queryResult.Columns {
			fmt.Printf("%-*s", colLengths[index], col)
			if index < colCount-1 {
				fmt.Print(options.separator)
			}
		}
		fmt.Println()

		for index := 0; index < colCount; index++ {
			fmt.Printf("%-*s", colLengths[index], strings.Repeat("-", colLengths[index]))
			if index < colCount-1 {
				fmt.Printf(options.separator)
			}
		}
		fmt.Println()
	}

	for _, row := range queryResult.Rows {
		for index, value := range row {
			fmt.Printf("%-*v", colLengths[index], *(value.(*interface{})))
			if index < colCount-1 {
				fmt.Print(options.separator)
			}
		}
		fmt.Println(options.lineSeparator)
	}
}
