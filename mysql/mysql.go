package mysql

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
	"toDoListProject/ToDoTask"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//represent the toDoTaskDB model
type toDoTaskDB struct {
	db *sql.DB
}

func openMySQLDB(dbUser, dbPass, ipAddr, portNo, dbName string) (*sql.DB, error) {
	return sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+ipAddr+":"+portNo+")/"+dbName)
}

func getDatabaseCredentials(jsonCredFile string) (string, string) {

	jsonFile, err := os.Open(jsonCredFile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened credentials.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var credentials Credentials

	err = json.Unmarshal(byteValue, &credentials)
	if err != nil {
		fmt.Println(err)
	}
	return credentials.Username, credentials.Password
}
func ConnectToDB(ipAddr, portNo, dbName string) (*sql.DB, error) {
	dbUser, dbPass := getDatabaseCredentials("credentials.json")
	db, err := openMySQLDB(dbUser, dbPass, ipAddr, portNo, dbName)
	if err != nil {
		fmt.Println("issa no openin'") // do something here
		log.Fatal(err)
	}
	return db, err
}

func (t *toDoTaskDB) InsertTask(toDoTaskObj *ToDoTask.ToDoTask) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	query := "INSERT INTO toDoListTest ( taskPriority, taskCheck," +
		"taskDescription, taskCategory, taskStartDate, taskDueDate)" +
		" VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := t.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, toDoTaskObj.TaskPriority, toDoTaskObj.TaskCheck,
		toDoTaskObj.TaskDescription, toDoTaskObj.TaskCategory,
		toDoTaskObj.TaskStartDate, toDoTaskObj.TaskDueDate)

	return err

}

// FindByID attaches the user repository and find data based on id
func (t *toDoTaskDB) FindByID(id int) (*ToDoTask.ToDoTask, error) {
	task := new(ToDoTask.ToDoTask)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := t.db.QueryRowContext(ctx,
		"SELECT taskID, taskCategory, taskCheck, taskDescription, taskPriority, "+
			"taskStartDate, taskDueDate"+
			" FROM toDoListTest WHERE id = ?", id).Scan(&task.TaskID, &task.TaskCategory,
		&task.TaskCheck, &task.TaskDescription, &task.TaskPriority, &task.TaskStartDate,
		&task.TaskDueDate)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// Find attaches the user repository and find all data
func (t *toDoTaskDB) Find() ([]*ToDoTask.ToDoTask, error) {
	tasks := make([]*ToDoTask.ToDoTask, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := t.db.QueryContext(ctx,
		"SELECT taskID, taskCategory, taskCheck, taskDescription, taskPriority, "+
			"taskStartDate, taskDueDate"+
			" FROM toDoListTest")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		task := new(ToDoTask.ToDoTask)
		err = rows.Scan(
			&task.TaskID,
			&task.TaskCategory,
			&task.TaskCheck,
			&task.TaskDescription,
			&task.TaskPriority,
			&task.TaskStartDate,
			&task.TaskDueDate,
		)

		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (toDoTaskDB *toDoTaskDB) Close() {
	toDoTaskDB.db.Close()
}
