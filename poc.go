package arin

import (
	"fmt"
	"strings"
)

type Contact struct {
	Name       string
	Handle     string
	Company    string
	Address1   string
	Address2   string
	City       string
	StateProv  string
	PostalCode string
	Country    string
	Registered string
	Updated    string
	Phone      string
	Email      string
}

func (c *Contact) addressString() string {
	var s []string

	s = append(s, c.Address1)

	if c.Address2 != "" {
		s = append(s, c.Address2)
	}

	s = append(s, fmt.Sprintf("%s, %s  %s", c.City, c.StateProv, c.PostalCode))
	s = append(s, c.Country)

	return strings.Join(s, "\n")
}

func (c *Contact) String() string {
	var s []string

	s = append(s, fmt.Sprintf("%s (%s)", c.Name, c.Handle))
	s = append(s, c.addressString())
	s = append(s, c.Phone)
	s = append(s, c.Email)
	s = append(s, fmt.Sprintf("Registered: %s", c.Registered))
	s = append(s, fmt.Sprintf("Updated: %s", c.Updated))

	return strings.Join(s, "\n")
}

func (c *Contact) Equal(c2 *Contact) bool {
	return c.Name == c2.Name &&
		c.Handle == c2.Handle &&
		c.Company == c2.Company &&
		c.Address1 == c2.Address1 &&
		c.Address2 == c2.Address2 &&
		c.City == c2.City &&
		c.StateProv == c2.StateProv &&
		c.PostalCode == c2.PostalCode &&
		c.Country == c2.Country &&
		c.Registered == c2.Registered &&
		c.Updated == c2.Updated &&
		c.Phone == c2.Phone &&
		c.Email == c2.Email
}

func NewContact(handle string) *Contact {
	var c = new(Contact)
	record := makeRequest("poc", handle)

	if record == "" {
		return c
	}

	rec := parseRecord(record)
	addr := strings.Split(rec["Address"], "\n")

	c.Name = rec["Name"]
	c.Handle = rec["Handle"]
	c.Company = rec["Company"]

	switch len(addr) {
	case 2:
		c.Address1 = addr[0]
		c.Address2 = addr[1]
	default:
		c.Address1 = addr[0]
	}

	c.City = rec["City"]
	c.StateProv = rec["StateProv"]
	c.PostalCode = rec["PostalCode"]
	c.Country = rec["Country"]
	c.Registered = rec["RegDate"]
	c.Updated = rec["Updated"]
	c.Phone = rec["Phone"]
	c.Email = rec["Email"]

	return c
}
