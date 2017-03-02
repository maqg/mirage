package main

import (
	"fmt"
	"octlink/mirage/src/utils/octmysql"
)

func main() {

	mysql := new(octmysql.OctMysql)

	rows, err := mysql.Query("SELECT ID,U_Name FROM tb_user WHERE U_Name=?", "root")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()

	var id string
	var name string
	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err == nil {
			fmt.Printf("query result: %s:%s\n", id, name)
		}
	}

}
