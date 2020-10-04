package main

import "fmt"

func main() {
	find := 13
	data := []int{2, 4, 5, 7, 9, 10, 13, 20, 21, 25, 30}

	fmt.Printf("index = %d\n\n", search(data, find))
	fmt.Printf("index = %d\n\n", searchRecursive(data, find, 0, len(data)-1))
}

func search(data []int, find int) int {
	low := 0
	high := len(data) - 1
	for low <= high {
		mid := low + ((high - low) / 2)
		fmt.Println(fmt.Sprintf("mid = %d, low = %d, high = %d", mid, low, high))
		if find == data[mid] {
			return mid
		}
		if find > data[mid] {
			low = mid + 1
		}
		if find < data[mid] {
			high = mid - 1
		}
	}
	return -1
}

func searchRecursive(data []int, find int, low int, high int) int {
	if low <= high {
		mid := low + ((high - low) / 2)
		fmt.Println(fmt.Sprintf("mid = %d, low = %d, high = %d", mid, low, high))
		if find == data[mid] {
			return mid
		}
		if find > data[mid] {
			low = mid + 1
		}
		if find < data[mid] {
			high = mid - 1
		}
		return searchRecursive(data, find, low, high)
	}
	return -1
}
