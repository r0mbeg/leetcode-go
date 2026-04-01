package p0344_reverse_string

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {

	testcase := []struct {
		str  []byte
		want []byte
	}{
		{
			[]byte{'h', 'e', 'l', 'l', 'o'},
			[]byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			[]byte{'h', 'a', 'n', 'n', 'a', 'h'},
			[]byte{'h', 'a', 'n', 'n', 'a', 'h'},
		},
	}

	for i, tc := range testcase {
		fmt.Printf("case %d init: %v ", i, tc.str)
		reverseString(tc.str)

		if !reflect.DeepEqual(tc.str, tc.want) {
			t.Fatalf("incorrect : got %v; expected: %v \n", tc.str, tc.want)
		}

		fmt.Printf("correct : got %v; expected: %v \n", tc.str, tc.want)
	}

}
