package ToDoTask

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
	TaskPriority    int
	TaskStartDate   time.Time
	TaskDueDate     string
}

func NewToDoTask(taskCategory, taskDescription string, taskPriority int, taskStartDate time.Time, taskDueDate string, taskCheck bool) *toDoTask {
	return &toDoTask{
		TaskCategory:    taskCategory,
		TaskCheck:       taskCheck,
		TaskDescription: taskDescription,
		TaskPriority:    taskPriority,
		TaskStartDate:   taskStartDate,
		TaskDueDate:     taskDueDate,
	}
}

func InsertTask(db *sql.DB, t *toDoTask) error {
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
		log.Printf("Error %s when inserting row into toDoListTest table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d toDoTasks created ", rows)
	return nil
}

//func multipleInsert(db *sql.DB, products []product) error {
//	query := "INSERT INTO product(product_name, product_price) VALUES "
//	var inserts []string
//	var params []interface{}
//	for _, v := range products {
//		inserts = append(inserts, "(?, ?)")
//		params = append(params, v.name, v.price)
//	}
//	queryVals := strings.Join(inserts, ",")
//	query = query + queryVals
//	log.Println("query is", query)
//	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancelfunc()
//	stmt, err := db.PrepareContext(ctx, query)
//	if err != nil {
//		log.Printf("Error %s when preparing SQL statement", err)
//		return err
//	}
//	defer stmt.Close()
//	res, err := stmt.ExecContext(ctx, params...)
//	if err != nil {
//		log.Printf("Error %s when inserting row into products table", err)
//		return err
//	}
//	rows, err := res.RowsAffected()
//	if err != nil {
//		log.Printf("Error %s when finding rows affected", err)
//		return err
//	}
//	log.Printf("%d products created simulatneously", rows)
//	return nil
//}

func removeTask(taskInputStr string) {

}
