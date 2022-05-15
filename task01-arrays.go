package homework

func average(input [15]float32) (result float32) {
	var sum, avg float32
	sum = 0

	for _, value := range input {
		sum = sum + value
	}

	avg = sum / float32(len(input))

	return avg
}
