package heatmap

import (
	"bytes"
	"fmt"
	"testing"
)

func TestHeatmapExamples_Examples(t *testing.T) {
	h := HeatmapExamples{}

	buf := bytes.NewBufferString("")
	h.Examples(buf)
	fmt.Printf("%v", buf)
}
