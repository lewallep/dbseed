package dbseed

import (
//	"context"
	"fmt"
	"database/sql"
	"math/rand"
	"time"

	//dsql "github.com/denisenkom/go-mssqldb"
)

func MssqlTest() {
	fmt.Println("dal_mssql.go has been called.")
}

func MssqlDropTableTdata(db *sql.DB) {
	var dropDb = `DROP DATABASE IF EXISTS tdata`

	_, dbErr := db.Exec(dropDb)

	if dbErr != nil {
		panic(dbErr)
	}
}

func MssqlCreateDbTdata(db *sql.DB) {
	createDb := `CREATE DATABASE tdata
	ON
	( NAME = Test_data,
		FILENAME = 'E:\databases\testdata.mdf',
		SIZE = 500000,
		MAXSIZE = 500000,
		FILEGROWTH = 100000 )
	LOG ON 
	( NAME = Test_log,
		FILENAME = 'E:\databases\testlog.ldf',
		SIZE = 500MB,
		MAXSIZE = 100000MB,
		FILEGROWTH = 100000MB );`

	_, dbErr := db.Exec(createDb)
	if dbErr != nil {
		panic(dbErr)
	}
}

func MssqlCreateTableUnitdata(db *sql.DB)  {
	dbt := `USE tdata
	CREATE TABLE dbo.unitdata
	(
		itemId int IDENTITY (1,1) NOT NULL
		,firstName VARCHAR(255)
		,lastName VARCHAR(255)
		,phoneNumber VARCHAR(13)
		,isbn10	VARCHAR(10)
		,isbn13 VARCHAR(13)
		,ccNum VARCHAR(16)
		,blurb VARCHAR(max)
		,email VARCHAR(255) NOT NULL
	);`

	_, dbErr := db.Exec(dbt)
	if dbErr != nil {
		panic(dbErr)
	}	
}

func (dal *Dal) MssqlGetCountUnitdata() int {
	var localRows int

	countRowsQ := `USE tdata SELECT COUNT(*) from unitdata;`
	rows, countErr := dal.Db.QueryContext(dal.Ctx, countRowsQ)
	if countErr != nil {
		panic(countErr)
	}
	for rows.Next() {
		err := rows.Scan(&localRows)
		if err != nil {
			panic(err)
		}
	}

	return localRows
}

var InsertOrder = `USE tdata
	INSERT INTO unitdata (
		firstName
		,lastName
		,phoneNumber
		,isbn10
		,isbn13
		,ccNum
		,blurb
		,email
	) values (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8)`

func (dal *Dal) ManualInsertUnitdata(fName string, lName string, phoneNum string, isbn10 int, 
		isbn13 int, ccNum int, blurb string, email string) {
	
	stmt, stmtErr := dal.Db.PrepareContext(dal.Ctx, InsertOrder)
	if stmtErr != nil {
		panic(stmtErr)
	}

	_, stmtErr = stmt.Exec(fName, lName, phoneNum, isbn10, isbn13, ccNum, blurb, email)
	_ = stmt.Close()
}

// Create a specified number of tables.
// Will create tables with a random number of columns specified by minCols and MaxCols
func (dal *Dal) CreateRandomTables(numTables int, minCols int, maxCols int, tablePrefix string) error {
	var ctSeg1 = `USE tdata CREATE TABLE dbo.` 
	var ctSeg2 = `(id int IDENTITY (1,1) NOT NULL`
	var ctSeg3 = `);`
	var r1 = rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < numTables; i++ {
		numCols := r1.Intn(maxCols) + minCols 	// Setting a random number for amount of columsn in the table.
		createTable := ctSeg1 + tablePrefix + fmt.Sprintf("%03d", i)
		createTable += ctSeg2

		// Inner loop to create the query for the desired amount of columns and type.
		for z := 0 ; z < numCols; z++ {
			createTable += randomCols(z)
		}

		createTable += ctSeg3

		stmt, stmtErr := dal.Db.PrepareContext(dal.Ctx, createTable)
		if stmtErr != nil {
			return stmtErr
		}
		_, stmtErr = stmt.ExecContext(dal.Ctx) 
		if stmtErr != nil {
			return stmtErr
		}
	}

	return nil
}

// Creates the string text of the column name and datatype
// Uses a pseudo random number generator for choosing the datatypes.
func randomCols(colNum int) string {
	var r1 = rand.New(rand.NewSource(time.Now().UnixNano()))
	var rn = r1.Intn(100)
	var colSeg = `, col_`

	colNum += 1
	switch {
	case rn < 60:
		return fmt.Sprintf("%s%03d VARCHAR(MAX)", colSeg, colNum)
	case rn >= 60 && rn < 70:
		return fmt.Sprintf("%s%03d INT", colSeg, colNum)
	case rn >= 70 && rn < 80:
		return fmt.Sprintf("%s%03d DATETIME2", colSeg, colNum)
	case rn >= 80 && rn < 90:
		return fmt.Sprintf("%s%03d MONEY", colSeg, colNum)
	default:
		return fmt.Sprintf("%s%03d VARCHAR(MAX)", colSeg, colNum)
	}
}

// Populates the tables created randomly with data
func (dal *Dal) InsertRandomData() error {
	// call function to get information about created tables
	// Return error if problematic.
	err := dal.getRandTableAndColMeta()
	if err != nil {
		return err
	}

	return nil
}

func (dal *Dal) getRandTableAndColMeta() error {
	var tableName, colId, colName, colDatatype string

	q := `USE tdata SELECT tab.name as table_name, 
    	col.column_id,
    	col.name as column_name, 
    	t.name as data_type  
	FROM sys.tables as tab
    INNER JOIN sys.columns as col
        on tab.object_id = col.object_id
    LEFT JOIN sys.types as t
    	on col.user_type_id = t.user_type_id
	ORDER BY table_name, 
    	column_id;`

    rows, err := dal.Db.Query(q)
    if err != nil {
    	return err
    }
    defer rows.Close()
    
    dal.tables = make(map[string]*TableMeta)

    for rows.Next() {
    	err = rows.Scan(&tableName, &colId, &colName, &colDatatype)
    	fmt.Printf("tableName: %s\n", tableName)
    	fmt.Printf("coldId: %s\n", colId)
    	fmt.Printf("colName: %s\n", colName)
    	fmt.Printf("colDatatype: %s\n", colDatatype)

    	if err != nil {
    		return err
    	}
    	// Check to see if the table name exists in the map.  If exists do nothing
    	if _, ok := dal.tables[tableName]; !ok {
    		dal.tables[tableName] = &TableMeta{}
    		dal.tables[tableName].cols = make(map[string]string)
    		dal.tables[tableName].cols[colName] = colDatatype
    	} else {
    		cols := dal.tables[tableName].cols
    		if _, ok := cols[colName]; !ok {
    			cols[colName] = colDatatype
    		}

    	}
    }

    fmt.Printf("tables len: %d\n", len(dal.tables))
    for k, v := range dal.tables {
    	fmt.Printf("key: %s\n", k)
    	fmt.Printf("cols: %v\n", v)
    }



	return err
}

