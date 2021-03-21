package mysql

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
	"toDoListProject/ToDoTask"
)

var tdTask = &ToDoTask.ToDoTask{
	TaskID:          0,
	TaskCategory:    "TestCategory",
	TaskCheck:       false,
	TaskDescription: "This is a task that needs to be completed",
	TaskPriority:    0,
	TaskStartDate:   time.Now().String(),
	TaskDueDate:     "2022-01-01 10:10:10",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}
func TestNewToDoTask(t *testing.T) {
	db, mock := NewMock()
	tdDB := &toDoTaskDB{db}
	defer func() {
		tdDB.Close()
	}()

	query := "INSERT INTO toDoListTest ( taskPriority, taskCheck," +
		"taskDescription, taskCategory, taskStartDate, taskDueDate)" +
		" VALUES (?, ?, ?, ?, ?, ?)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(tdTask.TaskPriority, tdTask.TaskCheck,
		tdTask.TaskDescription, tdTask.TaskCategory, tdTask.TaskStartDate,
		tdTask.TaskDueDate).WillReturnResult(sqlmock.NewResult(0, 1))

	err := tdDB.InsertTask(tdTask)
	assert.NoError(t, err)
}

func TestNewToDoTaskError(t *testing.T) {
	db, mock := NewMock()
	tdDB := &toDoTaskDB{db}
	defer func() {
		tdDB.Close()
	}()
	//the error is the table name
	query := "INSERT INTO toDoTest ( taskPriority, taskCheck," +
		"taskDescription, taskCategory, taskStartDate, taskDueDate)" +
		" VALUES (?, ?, ?, ?, ?, ?)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(tdTask.TaskPriority, tdTask.TaskCheck,
		tdTask.TaskDescription, tdTask.TaskCategory, tdTask.TaskStartDate,
		tdTask.TaskDueDate).WillReturnResult(sqlmock.NewResult(0, 0))

	err := tdDB.InsertTask(tdTask)
	assert.Error(t, err)
}

func TestFindByID(t *testing.T) {
	db, mock := NewMock()
	tdDB := &toDoTaskDB{db}
	defer func() {
		tdDB.Close()
	}()

	query := "SELECT taskID, taskCategory, taskCheck, taskDescription, taskPriority, " +
		"taskStartDate, taskDueDate" +
		" FROM toDoListTest WHERE id = ?"

	rows := sqlmock.NewRows([]string{"taskID", "taskCategory", "taskCheck", "taskDescription", "taskPriority",
		"taskStartDate", "taskDueDate"}).
		AddRow(tdTask.TaskID, tdTask.TaskCategory, tdTask.TaskCheck, tdTask.TaskDescription, tdTask.TaskPriority,
			tdTask.TaskStartDate, tdTask.TaskDueDate)

	mock.ExpectQuery(query).WithArgs(tdTask.TaskID).WillReturnRows(rows)

	task, err := tdDB.FindByID(tdTask.TaskID)
	assert.NotNil(t, task)
	assert.NoError(t, err)
}

func TestFindByIDError(t *testing.T) {
	db, mock := NewMock()
	tdDB := &toDoTaskDB{db}
	defer func() {
		tdDB.Close()
	}()
	//the error is the table name
	query := "SELECT taskID, taskCategory, taskCheck, taskDescription, taskPriority, " +
		"taskStartDate, taskDueDate" +
		" FROM test WHERE id = ?"

	rows := sqlmock.NewRows([]string{"taskID", "taskPriority", "taskCheck", "taskDescription", "taskCategory",
		"taskStartDate", "taskDueDate"})

	mock.ExpectQuery(query).WithArgs(tdTask.TaskID).WillReturnRows(rows)

	task, err := tdDB.FindByID(tdTask.TaskID)
	assert.Empty(t, task)
	assert.Error(t, err)
}

func TestFind(t *testing.T) {
	db, mock := NewMock()
	tdDB := &toDoTaskDB{db}
	defer func() {
		tdDB.Close()
	}()

	query := "SELECT taskID, taskCategory, taskCheck, taskDescription, taskPriority, " +
		"taskStartDate, taskDueDate" +
		" FROM toDoListTest"

	rows := sqlmock.NewRows([]string{"taskID", "taskPriority", "taskCheck", "taskDescription", "taskCategory",
		"taskStartDate", "taskDueDate"}).
		AddRow(tdTask.TaskID, tdTask.TaskCategory, tdTask.TaskCheck, tdTask.TaskDescription, tdTask.TaskPriority,
			tdTask.TaskStartDate, tdTask.TaskDueDate)

	mock.ExpectQuery(query).WillReturnRows(rows)

	tasks, err := tdDB.Find()
	assert.NotEmpty(t, tasks)
	assert.NoError(t, err)
	assert.Len(t, tasks, 1)
}

func TestUpdate(t *testing.T) {
	db, mock := NewMock()
	tdDB := &toDoTaskDB{db}
	defer func() {
		tdDB.Close()
	}()

	query := "UPDATE toDoListTest SET taskCategory = ?, taskCheck = ?, taskDescription = ?, " +
		"taskPriority = ?, taskStartDate = ?, taskDueDate = ? " +
		"WHERE taskID = ?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(tdTask.TaskCategory, tdTask.TaskCheck,
		tdTask.TaskDescription, tdTask.TaskPriority, tdTask.TaskStartDate,
		tdTask.TaskDueDate, tdTask.TaskID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := tdDB.Update(tdTask)
	assert.NoError(t, err)
}

func TestUpdateErr(t *testing.T) {
	db, mock := NewMock()
	tdDB := &toDoTaskDB{db}
	defer func() {
		tdDB.Close()
	}()

	//the error is the table name
	query := "UPDATE toDoTest SET taskCategory = ?, taskCheck = ?, taskDescription = ?, " +
		"taskPriority = ?, taskStartDate = ?, taskDueDate = ? " +
		"WHERE taskID = ?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(tdTask.TaskCategory, tdTask.TaskCheck,
		tdTask.TaskDescription, tdTask.TaskPriority, tdTask.TaskStartDate,
		tdTask.TaskDueDate, tdTask.TaskID).WillReturnResult(sqlmock.NewResult(0, 0))

	err := tdDB.Update(tdTask)
	assert.Error(t, err)
}

func TestDelete(t *testing.T) {
	db, mock := NewMock()
	tdDB := &toDoTaskDB{db}
	defer func() {
		tdDB.Close()
	}()

	query := "DELETE FROM toDoListTest WHERE taskID = ?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(tdTask.TaskID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := tdDB.Delete(tdTask.TaskID)
	assert.NoError(t, err)
}

func TestDeleteError(t *testing.T) {
	db, mock := NewMock()
	tdDB := &toDoTaskDB{db}
	defer func() {
		tdDB.Close()
	}()

	//the error is the table name
	query := "DELETE FROM toDoTest WHERE taskID = ?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(tdTask.TaskID).WillReturnResult(sqlmock.NewResult(0, 0))

	err := tdDB.Delete(tdTask.TaskID)
	assert.Error(t, err)
}
