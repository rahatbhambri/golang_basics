package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestComputeMerkle(t *testing.T) {
	a := []byte{
		252, 134, 128, 200, 6, 33, 31, 62,
		220, 18, 141, 8, 191, 132, 150, 126,
		101, 66, 184, 209, 5, 114, 119, 73,
		34, 103, 0, 5, 191, 174, 50, 53,
		102, 238, 237, 29, 28, 162, 250, 92,
		120, 29, 12, 59, 46, 18, 133, 87,
		18, 92, 89, 115, 65, 18, 118, 70,
		243, 10, 129, 140, 235, 108, 3, 191,
	}

	b := []byte{
		31, 165, 71, 133, 164, 136, 81, 51,
		3, 180, 158, 160, 187, 40, 129, 44,
		120, 47, 18, 152, 108, 174, 245, 253,
		94, 75, 73, 162, 76, 10, 168, 109,
		222, 62, 161, 223, 73, 254, 109, 110,
		218, 157, 52, 226, 86, 76, 181, 246,
		23, 54, 6, 152, 2, 206, 135, 217,
		87, 15, 164, 251, 23, 49, 206, 104,
	}

	h2 := simpleHash{"21.10.2023"}
	tests := []struct {
		name     string
		arr      []string
		expected []byte
	}{
		{"empty array", []string{}, nil},
		{"even length", []string{"Hi", "this", "i'm", "simple"}, a},
		{"odd length", []string{"Even", "Array"}, b},
	}

	for _, test := range tests {
		got := ComputeMerkle(test.arr, h2)

		if !bytes.Equal(got, test.expected) {
			t.Errorf("Expected %v but got %v", test.expected, got)
		} else {
			fmt.Printf("t: %v\n", t)
		}
	}
}
