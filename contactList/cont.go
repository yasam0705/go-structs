package contacts

type Contact struct {
	Id        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
}

type ContactList struct {
	Contacts []Contact
}

// Methods: create, update, get, getAll and delete

func (cl *ContactList) Create(con Contact) Contact {
	cl.Contacts = append(cl.Contacts, con)
	return con
}

func (cl *ContactList) Update(con Contact) Contact {
	for i, v := range cl.Contacts {
		if v.Id == con.Id {
			cl.Contacts[i] = con
			return cl.Contacts[i]
		}
	}
	return Contact{}
}

func (cl *ContactList) Get(id int) Contact {
	for _, v := range cl.Contacts {
		if v.Id == id {
			return v
		}
	}
	return Contact{}
}

func (cl *ContactList) GetAll() []Contact {
	return cl.Contacts
}

func (cl *ContactList) Delete(id int) []Contact {

	for i, v := range cl.Contacts {
		if v.Id == id {
			cl.Contacts = append(cl.Contacts[:i], cl.Contacts[i+1:]...)
			return cl.Contacts
		}
	}
	return cl.Contacts
}
