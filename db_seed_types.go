package dbseed

import (
	"context"
	"database/sql"
)

type Dal struct {
	Db 			*sql.DB
	Ctx 		context.Context
	tables 		map[string]*TableMeta	//Map of table names with a map of column names and type as the value
}

type TableMeta struct {
	cols 		map[string]string		// Key col name, value string col data type
}