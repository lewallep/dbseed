package dbseed

import (
	"context"
	"fmt"
	"database/sql"

	//dsql "github.com/denisenkom/go-mssqldb"
)

type Insert struct {
	Db 		*sql.DB
	Ctx 	context.Context
}

var (
	
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

func (dal *Insert) MssqlGetCountUnitdata() int {
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

func (dal *Insert) ManualInsertUnitdata(fName string, lName string, phoneNum string, isbn10 int, 
		isbn13 int, ccNum int, blurb string, email string) {
	
	stmt, stmtErr := dal.Db.PrepareContext(dal.Ctx, InsertOrder)
	if stmtErr != nil {
		panic(stmtErr)
	}

	_, stmtErr = stmt.Exec(fName, lName, phoneNum, isbn10, isbn13, ccNum, blurb, email)
	_ = stmt.Close()
}