package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	alex := Human{}
	alex.ID = 1
	alex.Name = "Alex"
	alex.Age = 25
	alex.Phone = "123-456-7890"
	alex.DateOfBirth = "1995-01-01"

	const file string = "humans.db"

	db, err := sql.Open("sqlite3", file)

	if err != nil {
		panic(err)
	}

	//res, err := db.Exec("INSERT INTO humans (id, name, age, phone, dateOfBirth) VALUES (?, ?, ?, ?, ?)", alex.ID, alex.Name, alex.Age, alex.Phone, alex.DateOfBirth)
	
	if err != nil {
		panic(err)
	}

	//fmt.Println(res)

	res1, err := db.Query("SELECT * FROM humans WHERE id = 1")

	for res1.Next() {
		var id int
		var name string
		var age int
		var phone string
		var dateOfBirth string

		err = res1.Scan(&id, &name, &age, &phone, &dateOfBirth)

		if err != nil {
			panic(err)
		}

		fmt.Println("ID:", id)
		fmt.Println("Name:", name)
		fmt.Println("Age:", age)
		fmt.Println("Phone:", phone)
		fmt.Println("Date of Birth:", dateOfBirth)

	}

	res1.Close()

}
