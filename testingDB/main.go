package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	PortHandler = 0
	Handler     = ""
	PathBD      = ""
	db          *sql.DB
)

func main() {
	sqlconnection("marlo")
	n := 10000
	for true {
		for i := 0; i < n; i++ {

			go theads()
		}
	}
}

func sqlconnection(namebd string) {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "1234",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: namebd,
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func theads() {
	for true {
		sqlconnection("marlo")
		var err error
		defer db.Close()

		// Create a table
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS test_table (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(50) NOT NULL, PRIMARY KEY (id));")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// Insert some data
		stmt, err := db.Prepare("INSERT INTO test_table (name) VALUES (?)")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer stmt.Close()

		for i := 0; i < 100; i++ {
			_, err = stmt.Exec("Name")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		rows, err := db.Query("SELECT * FROM test_table")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			var name string
			err = rows.Scan(&id, &name)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(id, name)
		}
	}
}
