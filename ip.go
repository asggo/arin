package arin

import (
    "encoding/xml"
    "fmt"
    "log"
)

type NetBlock struct {
    Start   string `xml:"startAddress"`
    End     string `xml:"endAddress"`
    CidrLen string `xml:"cidrLength"`
}

func (n *NetBlock) String() string {
    return fmt.Sprintf("%s - %s", n.Start, n.End)
}

func (n *NetBlock) Cidr() string {
    return fmt.Sprintf("%s/%s", n.Start, n.CidrLen)
}

type Organization struct {
    Name      string `xml:"name,attr"`
    Handle    string `xml:"handle,attr"`
    Reference string `xml:"orgRef"`
}

func (o *Organization) String() string {
    return fmt.Sprintf("Name: %s\nHandle: %s\n", o.Name, o.Handle)
}

type Parent struct {
    Name      string `xml:"name,attr"`
    Handle    string `xml:"handle,attr"`
    Reference string `xml:"parentNetRef"`
}

func (p *Parent) String() string {
    return fmt.Sprintf("Name: %s\nHandle: %s\n", p.Name, p.Handle)
}

type WhoisIP struct {
    Name         string `xml:"name"`
    Handle       string `xml:"handle"`
    Parent       *Parent `xml:"parentNetRef"`
    Organization *Organization `xml:"orgRef"`
    NetBlocks    []*NetBlock `xml:"netBlocks>netBlock"`
}

func (w *WhoisIP) String() string {
    s := fmt.Sprintf("Name: %s\nHandle: %s\n", w.Name, w.Handle)

    for _, n := range w.NetBlocks {
        s = s + fmt.Sprintf("Net Range: %v\n", n)
    }

    s = s + fmt.Sprintf("Organization:\n%v\n", w.Organization)
    s = s + fmt.Sprintf("Parent: \n%v", w.Parent)

    return s
}

func NewWhoisIP(addr string) *WhoisIP {
    var wip = new(WhoisIP)

    data := makeRequest("ip", addr)
    err := xml.Unmarshal(data, wip)
    if err != nil {
        log.Printf("Parse Error: %v\n", err)
        return wip
    }

    fmt.Println(wip)

    return wip
}
