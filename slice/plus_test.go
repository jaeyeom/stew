package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusStrings(t *testing.T) {

	s1 := []string{"one", "two"}
	s2 := []string{"three", "four"}

	all := PlusStrings(s1, s2)

	if assert.Equal(t, 4, len(all)) {
		assert.Equal(t, "one", all[0])
		assert.Equal(t, "two", all[1])
		assert.Equal(t, "three", all[2])
		assert.Equal(t, "four", all[3])
	}

}

func TestPlusStrings_ensureCopy(t *testing.T) {

	s1 := []string{"one", "two", "three", "four", "five"}
	s2 := []string{"six", "seven"}

	all := PlusStrings(s1[:2], s2)
	s1[0] = "eight"

	if assert.Equal(t, 4, len(all)) {
		assert.Equal(t, "one", all[0])
		assert.Equal(t, "two", all[1])
		assert.Equal(t, "six", all[2])
		assert.Equal(t, "seven", all[3])
	}

	if assert.Equal(t, 5, len(s1)) {
		assert.Equal(t, "eight", s1[0])
		assert.Equal(t, "two", s1[1])
		assert.Equal(t, "three", s1[2])
		assert.Equal(t, "four", s1[3])
		assert.Equal(t, "five", s1[4])
	}

}

func TestPlusStrings_ensureCopyNilPlus(t *testing.T) {

	s1 := []string{"one", "two", "three", "four", "five"}

	all := PlusStrings(s1[:2], nil)
	s1[0] = "eight"

	if assert.Equal(t, 2, len(all)) {
		assert.Equal(t, "one", all[0])
		assert.Equal(t, "two", all[1])
	}

	if assert.Equal(t, 5, len(s1)) {
		assert.Equal(t, "eight", s1[0])
		assert.Equal(t, "two", s1[1])
		assert.Equal(t, "three", s1[2])
		assert.Equal(t, "four", s1[3])
		assert.Equal(t, "five", s1[4])
	}

}
