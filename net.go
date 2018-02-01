package arin

import (
	"fmt"
	"strings"
)

type Network struct {
	Name         string
	Handle       string
	Start        string
	End          string
	Cidr         string
	Type         string
	Registered   string
	Updated      string
	Parent       string
	Organization string
	Children     []string
}

func (n *Network) String() string {
	var s []string

    s = append(s, "Whois Network Record")
    s = append(s, "====================")
	s = append(s, fmt.Sprintf("%s (%s)", n.Name, n.Handle))
    s = append(s, n.Organization)
	s = append(s, fmt.Sprintf("%s - %s (%s)", n.Start, n.End, n.Cidr))
	s = append(s, n.Type)
	s = append(s, fmt.Sprintf("%s (%s)", n.Registered, n.Updated))
    s = append(s, "")
    s = append(s, fmt.Sprintf("Parent:\n  %s", n.Parent))
    s = append(s, "Children:")

    for _, c := range n.Children {
        s = append(s, fmt.Sprintf("  %s", c))
    }

	return strings.Join(s, "\n")
}

func (n *Network) Equal(n2 *Network) bool {
	return n.Name == n2.Name &&
		n.Handle == n2.Handle &&
		n.Start == n2.Start &&
		n.End == n2.End &&
		n.Cidr == n2.Cidr &&
		n.Type == n2.Type &&
		n.Registered == n2.Registered &&
		n.Updated == n2.Updated &&
		n.Parent == n2.Parent &&
		n.Organization == n2.Organization
}

func NewNetwork(record string) *Network {
	var n = new(Network)

	if record == "" {
		return n
	}

	rec := parseRecord(record)
	rng := strings.Split(rec["NetRange"], " - ")
	parB := strings.Index(rec["Parent"], "(") + 1
	parE := strings.Index(rec["Parent"], ")")
	org := strings.Index(rec["Organization"], "(") - 1

	n.Name = rec["NetName"]
	n.Handle = rec["NetHandle"]
	n.Start = rng[0]
	n.End = rng[1]
	n.Cidr = rec["CIDR"]
	n.Type = rec["NetType"]
	n.Registered = rec["RegDate"]
	n.Updated = rec["Updated"]
	n.Parent = rec["Parent"][parB:parE]
	n.Organization = rec["Organization"][:org]

	childRec := makeSubRequest("net", n.Handle, "children")
	n.Children = parseChildren(childRec)

	return n
}

func NewNetworkIP(addr string) *Network {
	return NewNetwork(makeRequest("ip", addr))
}

func NewNetworkHandle(handle string) *Network {
	return NewNetwork(makeRequest("net", handle))
}
