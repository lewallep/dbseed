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