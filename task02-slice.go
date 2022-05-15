package homework

func reverse(input []int64) (result []int64) {

	ln := len(input)
	a := make([]int64, ln, ln)
	ln--
	for _, value := range input {
		a[ln] = value
		ln--
	}
	return a

}
