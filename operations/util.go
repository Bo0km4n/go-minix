package operations

import (
	"bytes"
	"fmt"
)

func getResult(idx int, ope, opeString string) string {
	return fmt.Sprintf("%04x: %-10s %s", idx, ope, opeString)
}

func getOrgOpe(args ...byte) string {
	var buffer bytes.Buffer
	for _, v := range args {
		s := fmt.Sprintf("%02x", v)
		buffer.WriteString(s)
	}
	return buffer.String()
}

func getOpeString(prefix string, args ...string) string {
	var buffer bytes.Buffer
	buffer.WriteString(prefix)
	buffer.WriteString(" ")

	for i, v := range args {
		buffer.WriteString(v)
		if i != len(args)-1 {
			buffer.WriteString(", ")
		}
	}
	return buffer.String()
}
