package main

import "fmt"

func main() {
	find := 21
	data := []int{2, 4, 5, 7, 9, 10, 13, 20, 21, 25, 30}

	fmt.Printf("index = %d\n\n", search(data, find))
	fmt.Printf("index = %d\n\n", searchRecursive(data, find, 0, len(data)-1))
}

func search(data []int, find int) int {
	low := 0
	high := len(data) - 1

	for low <= high && find >= data[low] && find <= data[high] {
		posFloat := float64(low) + ((float64(find) - float64(data[low])) * (float64(high) - float64(low)) / (float64(data[high]) - float64(data[low])))
		pos := int(posFloat)
		fmt.Println(fmt.Sprintf("pos = %d, low = %d, high = %d", pos, low, high))
		if data[pos] == find {
			return pos
		}
		if data[pos] > find {
			high = pos - 1
		}
		if data[pos] < find {
			low = pos + 1
		}
	}
	return -1
}

func searchRecursive(data []int, find int, low int, high int) int {
	if low <= high && find >= data[low] && find <= data[high] {
		posFloat := float64(low) + ((float64(find) - float64(data[low])) * (float64(high) - float64(low)) / (float64(data[high]) - float64(data[low])))
		pos := int(posFloat)
		fmt.Println(fmt.Sprintf("pos = %d, low = %d, high = %d", pos, low, high))
		if data[pos] == find {
			return pos
		}
		if data[pos] > find {
			high = pos - 1
		}
		if data[pos] < find {
			low = pos + 1
		}
		return searchRecursive(data, find, low, high)
	}
	return -1
}
