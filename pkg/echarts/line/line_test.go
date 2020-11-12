package line

import (
	"bytes"
	"fmt"
	"testing"
)

func TestLineExamples_Examples(t *testing.T) {
	e := LineExamples{}
	buf := bytes.NewBufferString("")
	e.Examples(buf)
	fmt.Printf("%v", buf)
}
