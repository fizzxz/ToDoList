package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"toDoListProject/Database"
)

func main() {

	dbUser, dbPass := Database.GetDatabaseCredentials("credentials.json")
	db, err := Database.OpenMySQLDB(dbUser, dbPass, "172.17.0.2", "3306", "mysqlTest")
	if err != nil {
		fmt.Println("issa no openin'") // do something here
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("issa no pinging") // do something here
	}
	defer db.Close()
	fmt.Println("issa :ok:")
}
