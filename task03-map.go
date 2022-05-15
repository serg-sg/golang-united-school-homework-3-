package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {

	b := make([]int, len(input))
	c := make([]string, len(input))

	i := 0
	for key, _ := range input {
		b[i] = key
		i++
	}

	sort.Ints(b)

	for index, value := range b {
		c[index] = input[value]
	}

	return c

}
