package main

import "fmt"

// time complexity = O(n)
// space complexity = O(1)
func main() {
	data := []int{2, 0, 4, 1, 8, 5, 3, 9, 6, 7}
	find := 8
	fmt.Println(data)
	fmt.Printf("index of %d = %d \n", find, search(find, data))
}

func search(find int, data []int) int {
	for index, v := range data {
		if v == find {
			return index
		}
	}
	return -1
}
