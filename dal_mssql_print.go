package dbseed

import (
	"fmt"
)

func (dal *Dal) PrintTableQueries() {
	for k, v := range dal.tables {
		fmt.Printf("table name: %s\n", k)
		fmt.Printf("table query: %s\n", v.query)
	}
}

func (dal *Dal) PrintTableRowsToAdd() {
	var totalRows = 0.0
	for tName, t := range dal.tables {
		fmt.Printf("table name: %s, \trowcount: %f\n", tName, t.rowsToAdd)
		totalRows += t.rowsToAdd
	}

	fmt.Printf("Sum of rows to add: %f\n", totalRows)
}

func (dal *Dal) TableCount() {
	fmt.Printf("")
}