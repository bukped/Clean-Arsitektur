package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "(localhost:3306)/test")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO users VALUES('3', 'Zaky', 'Muhammad')")
	// defer insert.Close()
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	// defer the close till after the main function has finished
	// executing
	// defer db.Close()

}

// https://www.youtube.com/watch?v=DWNozbk_fuk
// https://tutorialedge.net/golang/golang-mysql-tutorial/
