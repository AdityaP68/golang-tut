package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("omkie")

	var fruits = []string{"peach", "orange"}
	fmt.Println(fruits)
	fmt.Printf("its type is %T \n", fruits)

	fruits = append(fruits, "mango", "apple")
	fmt.Println(fruits)
	fruits = append(fruits[:3])
	fmt.Println(fruits)

	score := make([]int, 4)
	score[0] = 234
	score[1] = 235
	score[2] = 236
	score[3] = 237

	// memory reallocated when using append hence no out of bound error
	score = append(score, 4, 5, 6, 7, 8)
	sort.Ints(score)
	fmt.Println(score)
	fmt.Println(sort.IntsAreSorted(score))
}
