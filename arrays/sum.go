package arrays

// * Go has "slices" which do not encode the size of the collection and instead can have any size.

func Sum(numbers [5]int) int {
	// result := 0
	// for i := 0; i < len(numbers); i++ {
	// 	result += numbers[i]
	// }
	// return result

	result := 0
	// * for index, number := range numbers { ... }
	// * use "_" aka blank identifier since we are not using index value
	for _, number := range numbers {
		result += number
	}
	return result
}
