package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"toDoListProject/ToDoTask"
	"toDoListProject/mysql"
)

func main() {

	db, err := mysql.ConnectToDB("172.17.0.2", "3306", "mysqlTest")
	if err != nil {
		fmt.Println("Failed to open db")
		log.Fatal(err)
	}
	currTime := time.Now().String()
	//needs
	fmt.Print(currTime)

	dueTime := "2020-01-01 10:10:10"

	var newTask = ToDoTask.NewToDoTask(0, "", "Need To finish task",
		currTime, dueTime, 0, false)

	err = ToDoTask.TaskFunctions.InsertTask(db, newTask)
	if err != nil {
		fmt.Println("Failed to insert task into database")
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to ping") // do something here
	}

	defer db.Close()
}
