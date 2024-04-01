package dbconnect

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func LocalDbConnect() (sql_db *sql.DB, err error) {
	log.Println("LocalDbConnect (+)")
	driver := "mysql"
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "root", "192.168.2.5", 3306, "mani")
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		log.Println("DB CONNECTION FAILED :\n\n ", err.Error())
		return db, err
	} else {
		log.Println("LocalDbConnect (-)")
		return db, err
	}
}
