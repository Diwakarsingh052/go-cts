package sum

import "testing"

func TestSumInt(t *testing.T) { // Test prefix is compulsory followed by the funcName
	// t would be used to signal test failures
	x := []int{1, 2, 3, 4, 5}
	got := SumInt(x) // output from the func
	want := 15       // result expected

	if got != want {
		t.Errorf("sum of 1 to 5 should be %v; got %v", want, got) // signal test failure but continue to check other tests
		//t.Fatalf("sum of 1 to 5 should be %v; got %v", want, got) // stop the current test and exit
	}

	//some more things to do
	x = nil
	got = SumInt(x)
	want = 0
	if got != want {
		t.Errorf("sum of nil should be %v; got %v", want, got)
	}

	x = []int{1, -1}
	got = SumInt(x)
	want = 0
	if got != want {
		t.Errorf("sum of 1,-1 should be %v; got %v", want, got)
	}

}

// go test -v
//go test -cover
