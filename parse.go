package arin

import (
	"strings"
)

// getLines extracts the non-empty non-comment lines from the whois text record.
func getLines(record string) []string {
	var l []string

	for _, line := range strings.Split(record, "\n") {
		line = strings.Trim(line, "\r")

		if strings.HasPrefix(line, "#") {
			continue
		}
		if line == "" {
			continue
		}

		l = append(l, line)
	}

	return l
}

// parseRecord returns a map of keys and values from the whois text record.
// If there are duplicate keys, the values are concatenated into one string.
func parseRecord(record string) map[string]string {
	recMap := make(map[string]string)
	lines := getLines(record)

	for _, line := range lines {
		c := strings.Index(line, ":")
		k := line[:c]
		v := strings.Trim(line[c+1:], " ")

		// If the key is already in the map then append the data to the current
		// value.
		val, ok := recMap[k]
		switch {
		case ok:
			val = val + "\n" + v
			recMap[k] = val
		default:
			recMap[k] = v
		}
	}

	return recMap
}

// parseChildren returns a list of child network handles from the whois children
// record.
func parseChildren(record string) []string {
	var children []string

	lines := getLines(record)

	for _, line := range lines {
		s := strings.Index(line, "(") + 1
		e := strings.Index(line, ")")

		child := strings.Trim(line[s:e], " ")
		children = append(children, child)
	}

	return children
}
