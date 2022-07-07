package slices

func SumSlice(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

// // * here, we create a new slice after every append
// func SumAll(numbersToSum ...[]int) (sums []int) {
// 	for _, numbers := range numbersToSum {
// 		sum := 0
// 		for _, number := range numbers {
// 			sum += number
// 		}
// 		sums = append(sums, sum)
// 	}
// 	return

// 	// for _, numbers := range numbersToSum {
// 	// 	numbersSum := SumSlice(numbers)
// 	// 	sums = append(sums, numbersSum)
// 	// }
// 	// return
// }

// * here, we control the capacity of an array
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for index, numbers := range numbersToSum {
		sums[index] = SumSlice(numbers)
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for index, numbers := range numbersToSum {
		// * slice[low:high]
		// fmt.Println("")
		// fmt.Println(numbers)
		// fmt.Println("[:1]", numbers[:1])
		// fmt.Println("[:2]", numbers[:2])
		// fmt.Println("[1:2]", numbers[1:2]) // excludes low if low is given
		if len(numbers) == 0 {
			sums[index] = 0
		} else {
			sums[index] = SumSlice(numbers[1:])
		}
	}

	return sums
}

// * how slicing modifies original array but not copy of array
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	x := [3]string{"apple", "orange", "pear"}

// 	y := x[:] // slice "y" points to the underlying array "x"

// 	z := make([]string, len(x))
// 	copy(z, x[:]) // slice "z" is a copy of the slice created from array "x"

// 	y[1] = "Durian" // the value at index 1 is now "Belka" for both "y" and "x"

// 	fmt.Printf("x => %T %v\n", x, x)
// 	fmt.Printf("y => %T %v\n", y, y)
// 	fmt.Printf("z => %T %v\n", z, z)
// }

// * prints
// x => [3]string [apple Durian pear]
// y => []string [apple Durian pear]
// z => []string [apple orange pear]
