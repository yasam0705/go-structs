package tasklist

import (
	"testing"
)

var (
	testTasks = []Task{
		{
			Id:        10,
			Name:      "Task 10",
			Status:    "done",
			Priority:  "Important",
			CreatedAt: "04.10.2021",
			CreatedBy: "05.10.2021",
			DueDate:   "08.10.2021",
		},
		{
			Id:        11,
			Name:      "Task 11",
			Status:    "testing",
			Priority:  "medium",
			CreatedAt: "04.10.2021",
			CreatedBy: "",
			DueDate:   "09.10.2021",
		},
		{
			Id:        12,
			Name:      "Task 12",
			Status:    "Initial",
			Priority:  "medium",
			CreatedAt: "04.10.2021",
			CreatedBy: "",
			DueDate:   "10.10.2021",
		},
	}

	updTask = Task{
		Id:        11,
		Name:      "Task 10",
		Status:    "done",
		Priority:  "Important",
		CreatedAt: "04.10.2021",
		CreatedBy: "06.10.2021",
		DueDate:   "09.10.2021",
	}
	TaskL = TaskList{}
)

func TestCreate(t *testing.T) {
	for _, v := range testTasks {
		temp_task := TaskL.Create(v)
		if temp_task != v {
			t.Error("Failed create method")
		}
	}

	t.Cleanup(func() {
		TaskL = TaskList{}
	})
}

func TestUpdate(t *testing.T) {
	TestCreate(t)

	TaskL.Update(updTask)
	if TaskL.Get(updTask.Id) != updTask {
		t.Error("Failed Update method")
	}
}

func TestGet(t *testing.T) {
	TestCreate(t)

	getTask := TaskL.Get(12)
	if getTask.Id != 12 {
		t.Error("Failed get method")
	}
}

func TestGetAll(t *testing.T) {
	TestCreate(t)
	allTasks := TaskL.GetAll()

	for i := 0; i < len(testTasks); i++ {
		if allTasks[i] != testTasks[i] {
			t.Error("Failed GetAll method")
		}
	}
}

func TestDelete(t *testing.T) {
	TestCreate(t)
	TaskL.Delete(10)
	for _, v := range TaskL.Tasks {
		if v.Id == 10 {
			t.Error("Failed Delete method")
		}
	}
}
