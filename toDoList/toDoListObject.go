package toDoList

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type toDoTask struct {
	TaskID          int
	TaskCategory    string
	TaskCheck       bool
	TaskDescription string
	TaskPriority    string
	TaskDueDate     string
	TaskStartDate   string
}

func NewToDoList(taskCategory, taskDescription, taskPriority, taskStartDate, taskDueDate string, taskCheck bool) *toDoTask {
	return &toDoTask{
		TaskCategory:    taskCategory,
		TaskCheck:       taskCheck,
		TaskDescription: taskDescription,
		TaskPriority:    taskPriority,
		TaskDueDate:     taskDueDate,
		TaskStartDate:   taskStartDate,
	}
}

func insert(db *sql.DB, t toDoTask) error {
	query := "INSERT INTO toDoListTest( taskPriority, taskCheck," +
		"taskDescription, taskCategory, taskStartDate, taskDueDate)" +
		" VALUES (?, ?, ?, ?, ?, ?)"
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, t.TaskPriority, t.TaskCheck, t.TaskDescription,
		t.TaskCategory, t.TaskStartDate, t.TaskDueDate)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d products created ", rows)
	return nil
}

func removeTask(taskInputStr string) {

}
