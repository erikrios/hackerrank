package main

import "testing"

func equalizeArrayTest(t *testing.T) {
	testCases := []struct {
		name string
		arr  []int32
		want int32
	}{
		{
			name: "First case",
			arr:  []int32{1, 2, 2, 3},
			want: 2,
		},
		{
			name: "Second case",
			arr:  []int32{3, 3, 2, 1, 2},
			want: 2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := equalizeArray(testCase.arr)
			if got != testCase.want {
				t.Errorf("want %d, but got %d", testCase.want, got)
			}
		})
	}
}
