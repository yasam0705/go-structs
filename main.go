package main

import (
	contacts "contact_list/contactList"
	tasklist "contact_list/taskList"
)

func main() {
	cont := new(contacts.ContactList)

	cont.Create(contacts.Contact{
		Id:        10,
		FirstName: "Sam",
		LastName:  "Smith",
		Phone:     "991234567",
		Email:     "sam@local.com",
	})
	cont.Create(contacts.Contact{
		Id:        12,
		FirstName: "Eugene",
		LastName:  "Williamson",
		Phone:     "(139)-191-0039",
		Email:     "eugene@local.edit",
	})

	// fmt.Println(cont.GetAll())

	// cont.Delete(12)

	// fmt.Println(cont.Get(12))

	taskl := new(tasklist.TaskList)

	taskl.Create(tasklist.Task{
		Id:        15,
		Name:      "task15",
		Status:    "asd",
		Priority:  "medium",
		CreatedAt: "03.10.2021",
		CreatedBy: "03.10.2021",
		DueDate:   "05.10.2021",
	})

	taskl.Create(tasklist.Task{
		Id:        20,
		Name:      "task20",
		Status:    "qwerty",
		Priority:  "important",
		CreatedAt: "03.10.2021",
		CreatedBy: "03.10.2021",
		DueDate:   "06.10.2021",
	})

	// taskl.Delete(20)
	// fmt.Println(taskl.GetAll())
}
