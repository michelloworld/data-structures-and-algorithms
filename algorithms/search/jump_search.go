package main

import (
	"fmt"
	"math"
)

func main() {
	find := 25
	data := []int{2, 4, 5, 7, 9, 10, 13, 20, 21, 25, 30, 31, 36, 38, 45, 55}

	fmt.Printf("index = %d\n", search(data, find))
	fmt.Println("---")
	fmt.Printf("index = %d\n", search2(data, find))
	fmt.Println("---")
	fmt.Printf("index = %d\n", searchRecursive(data, find, 0, int(math.Sqrt(float64(len(data))))))
}

func search(data []int, find int) int {
	left := 0
	right := int(math.Sqrt(float64(len(data)))) - 1
	if data[left] == find {
		return left
	}
	if data[right] == find {
		return right
	}
	for data[right] < find && right < len(data) {
		left = right
		right += int(math.Sqrt(float64(len(data)))) - 1
		fmt.Println(fmt.Sprintf("left = %d, right = %d", left, right))
		if data[left] == find {
			return left
		}
		if data[right] == find {
			return right
		}
	}
	for i := left + 1; i < right; i++ {
		fmt.Println(fmt.Sprintf("left = %d, right = %d, i = %d", left, right, i))
		if data[i] == find {
			return i
		}
	}
	return -1
}

func search2(data []int, find int) int {
	left := 0
	right := int(math.Sqrt(float64(len(data))))
	for right < len(data) && data[right] <= find {
		fmt.Println(fmt.Sprintf("left = %d, right = %d", left, right))
		left = right
		right += int(math.Sqrt(float64(len(data))))
		if right > len(data)-1 {
			right = len(data)
		}
	}
	for i := left; i < right; i++ {
		fmt.Println(fmt.Sprintf("left = %d, right = %d, i = %d", left, right, i))
		if data[i] == find {
			return i
		}
	}
	return -1
}

func searchRecursive(data []int, find int, left int, right int) int {
	if right < len(data) && data[right] <= find {
		fmt.Println(fmt.Sprintf("left = %d, right = %d", left, right))
		left = right
		right += int(math.Sqrt(float64(len(data))))
		if right > len(data)-1 {
			right = len(data)
		}
		return searchRecursive(data, find, left, right)
	}
	for i := left; i < right; i++ {
		fmt.Println(fmt.Sprintf("left = %d, right = %d, i = %d", left, right, i))
		if data[i] == find {
			return i
		}
	}
	return -1
}
