package arin

import (
    "strings"
    "fmt"
)

type Contact struct {
    Name       string
    Handle     string
    Company    string
    Address    string
    City       string
    StateProv  string
    PostalCode string
    Country    string
    Registered string
    Updated    string
    Phone      string
    Email      string
}

func (c *Contact) String() string {
    var s []string

    s = append(s, fmt.Sprintf("Name: %s", c.Name))
    s = append(s, fmt.Sprintf("Handle: %s", c.Handle))
    s = append(s, fmt.Sprintf("Company: %s", c.Company))
    s = append(s, fmt.Sprintf("Address: %s", c.Address))
    s = append(s, fmt.Sprintf("City: %s", c.City))
    s = append(s, fmt.Sprintf("StateProv: %s", c.StateProv))
    s = append(s, fmt.Sprintf("PostalCode: %s", c.PostalCode))
    s = append(s, fmt.Sprintf("Registered: %s", c.Registered))
    s = append(s, fmt.Sprintf("Updated: %s", c.Updated))
    s = append(s, fmt.Sprintf("Phone: %s", c.Phone))
    s = append(s, fmt.Sprintf("Email: %s", c.Email))

    return strings.Join(s, "\n")
}

func (c *Contact) Equal(c2 *Contact) bool {
    return c.Name == c2.Name &&
        c.Handle == c2.Handle &&
        c.Company == c2.Company &&
        c.Address == c2.Address //&&
        // c.City == c2.City &&
        // c.StateProv == c2.StateProv &&
        // c.PostalCode == c2.PostalCode &&
        // c.Registered == c2.Registered &&
        // c.Updated == c2.Updated &&
        // c.Phone == c2.Phone &&
        // c.Email == c2.Email
}

func NewContact(handle string) *Contact {
    var c = new(Contact)
    record := makeRequest("poc", handle)

    if record == "" {
        return c
    }

    rec := parseRecord(record)

    c.Name = rec["Name"]
    c.Handle = rec["Handle"]
    c.Company = rec["Company"]
    c.Address = rec["Address"]
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
