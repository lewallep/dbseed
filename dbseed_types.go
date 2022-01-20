package dbseed

import (
	"context"
	"database/sql"
)

type Dal struct {
	Db 			*sql.DB
	Ctx 		context.Context
	tables 		map[string]*TableMeta	//Map of table names with a map of column names and type as the value
	RowsToAdd	int 					// Total rows to add for the run.
	NumTables   int
}

type TableMeta struct {
	cols 		map[string]string		// Key col name, value string col data type
	colsAsc		[]string				// Sorted in Ascending order.
	query		string
	rowsToAdd	float64					// How many rows to add during this specific run of the program.
}
