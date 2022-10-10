package nicerslicer

import (
	"strings"

	"github.com/srfrog/slices"
)

type NicerSlicer struct {
	slices []string
}

func NewNicerSlicer(sentence string) *NicerSlicer {
	 result NicerSlicer = NicerSlicer{
		slices: strings.Split(sentence, " "),
	}
	return &result
}

func NewNicerSlicerFromSlices(sentence []string) *NicerSlicer {
	var result NicerSlicer = NicerSlicer{
		slices: sentence,
	}
	return &result
}

func (nc *NicerSlicer) CountWord(word string) int {
	return slices.Count(nc.slices, word)
}

func (this *NicerSlicer) Split(sep string) []NicerSlicer {
	var result []NicerSlicer

	var newSlices = slices.Split(this.slices, sep)
	for _, slice := range newSlices {
		result = append(result, *NewNicerSlicerFromSlices(slice))
	}

	return result
}

func (nc *NicerSlicer) Chunk(size int) [][]string {
	// Chunk the slice
	chunks := slices.Chunk(nc.slices, size)
	return chunks
}
