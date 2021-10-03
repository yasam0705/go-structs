package tasklist

type Task struct {
	Id        int
	Name      string
	Status    string
	Priority  string
	CreatedAt string
	CreatedBy string
	DueDate   string
}

type TaskList struct {
	Tasks []Task
}

// Methods: create, update, get, getAll and delete

func (tl *TaskList) Create(t Task) Task {
	tl.Tasks = append(tl.Tasks, t)
	return t
}

func (tl *TaskList) Update(t Task) Task {
	for i, v := range tl.Tasks {
		if v.Id == t.Id {
			tl.Tasks[i] = t
			return t
		}
	}
	return Task{}
}

func (tl *TaskList) Get(id int) Task {
	for _, v := range tl.Tasks {
		if v.Id == id {
			return v
		}
	}
	return Task{}
}

func (tl *TaskList) GetAll() []Task {
	return tl.Tasks
}

func (tl *TaskList) Delete(id int) []Task {
	for i, v := range tl.Tasks {
		if v.Id == id {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			return tl.Tasks
		}
	}
	return tl.Tasks
}
