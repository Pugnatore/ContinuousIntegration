package nicerslicer

import (
	"testing"
)

func TestSplit(t *testing.T) {
	str := `Don't communicate by sharing memory - share memory by communicating`
	sl := NewNicerSlicer(str)

	splitSlices := sl.Split("-")

	if len(splitSlices) != 2 {
		t.Errorf(`Expected slice length of 2. Got %v`, len(splitSlices))
	}
}

func TestChunk(t *testing.T) {
	str := `Don't communicate by sharing memory`
	sl := NewNicerSlicer(str)

	splitChunks := sl.Chunk(2)

	if len(splitChunks) != 3 {
		t.Fatalf(`Expected 3 chunks. Got %v`, len(splitChunks))
	}

	if len(splitChunks[0]) != 2 || len(splitChunks[1]) != 2 {
		t.Fatalf(`Expected all chunks to be of length 2 (except for the last one). Got %v`, splitChunks)
	}

	if len(splitChunks[2]) != 1 {
		t.Fatalf(`Expected last chunk to be of length 1. Got %v`, splitChunks)
	}
}
