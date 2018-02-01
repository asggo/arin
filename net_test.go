package arin

import (
	"testing"
)

var n = &Network{
	Name:         "AT-88-Z",
	Handle:       "NET-52-32-0-0-1",
	Start:        "52.32.0.0",
	End:          "52.63.255.255",
	Cidr:         "52.32.0.0/11",
	Type:         "Direct Allocation",
	Registered:   "2015-09-02",
	Updated:      "2015-09-02",
	Parent:       "NET-52-0-0-0-0",
	Organization: "Amazon Technologies Inc.",
}

func TestNewNetworkIP(t *testing.T) {
	i := NewNetworkIP("52.52.52.0")
	h := NewNetworkHandle("NET-52-32-0-0-1")

	if !i.Equal(n) {
		t.Errorf("Expected\n%s\ngot\n%s\n", n, i)
	}

	if !h.Equal(n) {
		t.Errorf("Expected\n%s\ngot\n%s\n", n, h)
	}
}
