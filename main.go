package main

import (
	driver "arjun/infotainment/drivers"
	"database/sql"
	"fmt"
	"log"

	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Println(err)
	}
}

func main() {

	db = driver.ConnectDB()

	fmt.Println("Hello infotainment...")
}
