package ToDoTask

type ToDoTask struct {
	TaskID          int
	TaskCategory    string
	TaskCheck       bool
	TaskDescription string
	TaskPriority    int
	TaskStartDate   string
	TaskDueDate     string
}

func NewToDoTask(taskID int, taskCategory, taskDescription, taskStartDate,
	taskDueDate string, taskPriority int, taskCheck bool) *ToDoTask {
	return &ToDoTask{
		TaskID:          taskID,
		TaskCategory:    taskCategory,
		TaskCheck:       taskCheck,
		TaskDescription: taskDescription,
		TaskPriority:    taskPriority,
		TaskStartDate:   taskStartDate,
		TaskDueDate:     taskDueDate,
	}
}

type TaskFunctions interface {
	Close()
	FindByID(id string) (*ToDoTask, error)
	Find() ([]*ToDoTask, error)
	InsertTask(user *ToDoTask) error
	Update(user *ToDoTask) error
	Delete(id string) error
}
