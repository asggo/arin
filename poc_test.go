package arin

import (
	"testing"
)

var c = &Contact{
	Name:       "Registration Services Department",
	Handle:     "ARIN-HOSTMASTER",
	Company:    "American Registry for Internet Numbers",
	Address1:   "3635 Concorde Parkway",
	Address2:   "Suite 200",
	City:       "Chantilly",
	StateProv:  "VA",
	PostalCode: "20121",
	Country:    "US",
	Registered: "2003-04-30",
	Updated:    "2011-07-21",
	Phone:      "+1-703-227-0660 (Office)",
	Email:      "hostmaster@arin.net",
}

func TestNewContact(t *testing.T) {
	contact := NewContact("ARIN-HOSTMASTER")

	if !contact.Equal(c) {
		t.Errorf("Expected\n%s\ngot\n%s\n", c, contact)
	}
}
