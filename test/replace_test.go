package test

import (
	"fmt"
	"regexp"
	"testing"
)

func TestReplace(t *testing.T) {
	text := "<!-- BEGIN -->123<!-- END -->\n 456"
	reg,_ := regexp.Compile(`<!-- BEGIN -->[\S\s]+<!-- END -->`)
	allString := reg.ReplaceAllString(text, "<!-- BEGIN -->666<!-- END -->")
	fmt.Println(allString)
}
