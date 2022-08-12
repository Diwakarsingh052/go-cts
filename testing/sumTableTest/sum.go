package sum

func SumInt(vs []int) int {

	sum := 0
	if vs == nil {
		return 0 // we test the return statements
	}
	for _, v := range vs {
		sum = v + sum
	}
	return sum // we test the return statements

}
