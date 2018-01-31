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

func (n *NetBlock) Cidr() string {
    return fmt.Sprintf("%s/%s", n.Start, n.CidrLen)
}

type Organization struct {
    Name      string `xml:"name,attr"`
    Reference string `xml:"orgRef"`
}

type Parent struct {
    Name      string `xml:"name,attr"`
    Handle    string `xml:"handle,attr"`
    Reference string `xml:"parentNetRef"`
}

type WhoisIP struct {
    Name         string `xml:"name"`
    Handle       string `xml:"handle"`
    Parent       *Parent `xml:"parentNetRef"`
    Organization *Organization `xml:"orgRef"`
    Netblocks    []*NetBlock `xml:"netBlocks>netBlock"`
}

func GetWhoisIP(addr string) *WhoisIP {
    var wip = new(WhoisIP)

    data := makeRequest("ip", addr)
    err := xml.Unmarshal(data, wip)
    if err != nil {
        log.Printf("Parse Error: %v\n", err)
        return wip
    }

    return wip
}
