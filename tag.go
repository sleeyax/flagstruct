package flag

import (
	"strings"
)

const tagName = "flag"

// ParseTag parses a struct tag into a Flag.
func ParseTag(tag string) *Flag {
	flag := &Flag{}

	for _, pair := range strings.Split(tag, ",") {
		kvp := strings.Split(pair, "=")
		if len(kvp) < 2 {
			continue
		}
		key := strings.TrimSpace(kvp[0])
		value := strings.TrimSpace(kvp[1])

		switch key {
		case "value":
			flag.Value = value
		case "usage":
			flag.Usage = value
		case "name":
			flag.Name = value
		}
	}

	return flag
}
