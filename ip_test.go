package arin

import (
    "testing"
)


func TestGetWhoisIP(t *testing.T) {
    orgName := "Amazon Technologies Inc."
    ip := GetWhoisIP("52.52.52.0")

    if ip == nil {
        t.Error("Expected an IP record but got nil.")
    }

    if ip.Organization.Name != orgName {
        t.Error("Expected", orgName, "got", ip.Parent.Name)
    }
}
