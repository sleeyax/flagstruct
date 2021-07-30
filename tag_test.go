package flag

import (
	"fmt"
	"github.com/sleeyax/flagstruct/internal/tests"
	"testing"
)

func TestParseTag(t *testing.T) {
	name := "foo"
	value := "bar"
	usage := "foo contains bar"
	tag := fmt.Sprintf("name=%s,value=%s,usage=%s", name, value, usage)

	for i := 0; i < 2; i++ {
		flag := ParseTag(tag)

		if flag.Name != name {
			t.Errorf(tests.MismatchFormat, name, flag.Name)
		}
		if flag.Value != value {
			t.Errorf(tests.MismatchFormat, value, flag.Value)
		}
		if flag.Usage != usage {
			t.Errorf(tests.MismatchFormat, usage, flag.Usage)
		}

		if t.Failed() {
			fmt.Println(fmt.Sprintf("loop index: %d", i))
		}

		// test again with weird formatting
		if i == 0 {
			tag = fmt.Sprintf(",,  name =   %s , value =%s, usage= %s", name, value, usage)
		}

		// test again, but this time with the optional 'value' omitted
		if i == 1 {
			tag = fmt.Sprintf("value=%s,usage=%s", value, usage)
			name = ""
		}
	}
}
