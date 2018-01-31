package arin

import (
    "testing"
)


func TestNewWhoisIP(t *testing.T) {
    name := "AT-88-Z"
    handle := "NET-52-32-0-0-1"
    parentName := "NET52"
    parentHandle := "NET-52-0-0-0-0"
    orgName := "Amazon Technologies Inc."
    orgHandle := "AT-88-Z"
    cidr := "52.32.0.0/11"

    ip := NewWhoisIP("52.52.52.0")

    if ip == nil {
        t.Error("Expected an IP record but got nil.")
    }

    if ip.Name != name {
        t.Error("Expected", name, "got", ip.Name)
    }

    if ip.Handle != handle {
        t.Error("Expected", handle, "got", ip.Handle)
    }

    if ip.Parent.Name != parentName {
        t.Error("Expected", parentName, "got", ip.Parent.Name)
    }

    if ip.Parent.Handle != parentHandle {
        t.Error("Expected", parentHandle, "got", ip.Parent.Handle)
    }

    if ip.Organization.Name != orgName {
        t.Error("Expected", orgName, "got", ip.Organization.Name)
    }

    if ip.Organization.Handle != orgHandle {
        t.Error("Expected", orgHandle, "got", ip.Organization.Handle)
    }

    if ip.NetBlocks[0].Cidr() != cidr {
        t.Error("Expected", cidr, "got", ip.NetBlocks[0].Cidr())
    }
}
