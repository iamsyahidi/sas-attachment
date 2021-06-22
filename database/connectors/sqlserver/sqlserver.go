package sqlserver

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
)


var (
	dbMaxIdleConns = 4
	dbMaxConns     = 100
)

// Connect to db func
func ConnectDb() (*sql.DB, error) {

	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DBNAME")

	conString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", host, user, password, port, dbname)

	db, err := sql.Open("sqlserver", conString)
	if err != nil {
		fmt.Println("Failed to connect database!", err)
		return nil, err
	}

	ctx := context.Background()

	testPing := db.PingContext(ctx)
	if testPing != nil {
		log.Fatal("Error pinging database: " + testPing.Error())
	}

	db.SetMaxOpenConns(dbMaxConns)
	db.SetMaxIdleConns(dbMaxIdleConns)

	return db, nil
}

// Query func
func Query(sql string) (*sql.Rows, error) {
	ctx := context.Background()
	db, _ := ConnectDb()
	defer db.Close()
	return db.QueryContext(ctx, sql)
}

// QueryCount func
func QueryCount(queryCount string) int {
	var totalRows int
	rowsCount, _ := Query(queryCount)
	for rowsCount.Next() {
		errScanCount := rowsCount.Scan(&totalRows)
		if errScanCount != nil {
			return 0
		}
	}
	return totalRows
}