package contacts

import "testing"

var (
	testContacts = []Contact{
		{
			Id:        10,
			FirstName: "Sam",
			LastName:  "Smith",
			Phone:     "(695)-175-4661",
			Email:     "sam@local.com",
		},
		{
			Id:        11,
			FirstName: "Eugene",
			LastName:  "Williamson",
			Phone:     "(139)-191-0039",
			Email:     "eugene@local.com",
		},
		{
			Id:        12,
			FirstName: "Brian",
			LastName:  "Robinson",
			Phone:     "(045)-207-9455",
			Email:     "brian.robinson@example.com",
		},
	}
	contList = ContactList{}
)

func TestCreate(t *testing.T) {
	t.Run("create contact", func(t *testing.T) {
		for _, v := range testContacts {
			temp_contact := contList.Create(v)
			if temp_contact != v {
				t.Errorf("Failed create object: %v", v)
			}
		}
	})
}

func TestUpdate(t *testing.T) {
	TestCreate(t)
}

func TestGet(t *testing.T) {
	TestCreate(t)
	getContact := contList.Get(11)
	if getContact.Id != 11 {
		t.Error("Failed get method")
	}
}

func TestGetAll(t *testing.T) {
	TestCreate(t)
	allContacts := contList.GetAll()
	// if allContacts == testContacts {
	// 	t.Error("Failed GetAll method")
	// }
}
