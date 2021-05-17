package proch

import (
	"testing"
)

func TestLoad(t *testing.T) {
	jl := NewJsonLoader()

	testCases := []struct{
		filepath string
		expected int
	}{
		{
			"./test/test.json",
			2,
		},
	}

	for i, tc := range testCases {
		result, err := jl.Load(tc.filepath)
		if err != nil {
			t.Errorf("test case %d: failed to load %s", i, tc.filepath)
		}
		if len(result) != tc.expected {
			t.Errorf("test case %d: expected %d, result %d", i, tc.expected, len(result))
		}
	}
}