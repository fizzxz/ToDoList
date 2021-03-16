package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"toDoListProject/Database"
	"toDoListProject/ToDoTask"
)

func main() {

	db, err := Database.ConnectToDB("172.17.0.2", "3306", "mysqlTest")
	if err != nil {
		fmt.Println("issa no openin'") // do something here
		log.Fatal(err)
	}
	currTime := time.Now()
	//needs
	fmt.Print(currTime)

	dueTime := "2020-01-01 10:10:10"
	newTask := ToDoTask.NewToDoTask("", "Need To finish task", 1, currTime, dueTime, false)
	_ = ToDoTask.InsertTask(db, newTask)
	err = db.Ping()
	if err != nil {
		fmt.Println("issa no pinging") // do something here
	}

	defer db.Close()
	fmt.Println("issa :ok:")
}
