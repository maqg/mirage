package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/dbmirage?charset=utf8")
	if err != nil {
		fmt.Println("failed to open database:", err.Error())
		return
	}
	defer db.Close()

	fmt.Printf("Connect OK!\n")

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT ID,U_Name FROM tb_user WHERE Id = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	var id string
	var name string

	err = db.QueryRow("SELECT ID,U_Name FROM tb_user WHERE ID='8cd448b4c15543808c48932b152bad42'").Scan(&id, &name) // WHERE number = 13
	//err = stmtOut.QueryRow("8cd448b4c15543808c48932b152bad42").Scan(&id, &name) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("query result: %s:%s\n", id, name)

	rows, err := db.Query("SELECT ID,U_Name FROM tb_user")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err == nil {
			fmt.Printf("query result: %s:%s\n", id, name)
		}
	}

}
