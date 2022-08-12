package sum

import "testing"

/*

	name	 age	 marks
    abc		 6  		78
	xyz 	  22		77

*/

func TestSumInt(t *testing.T) {
	tt := []struct {
		name    string
		numbers []int
		want    int
	}{
		{
			name:    "one to five",
			numbers: []int{1, 2, 3, 4, 5}, //test case
			want:    15,
		},
		{
			name:    "nil slice",
			numbers: nil,
			want:    0,
		},
		{
			name:    "one minus one",
			numbers: []int{1, -1},
			want:    0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) { // t.Run() creates the subtest which will run test cases as individual tests
			got := SumInt(tc.numbers)
			if got != tc.want {
				t.Errorf("sum of %v want %v; got %v", tc.numbers, tc.want, got)
			}
		})
	}

}

//go test -run TestSumInt/one -v

func TestAnyName(t *testing.T) {

}
