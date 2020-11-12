package gauge

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGaugeExamples_Examples(t *testing.T) {
	g := GaugeExamples{}
	buf := bytes.NewBufferString("")

	g.Examples(buf)
	fmt.Printf("%v", buf)
}
