package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func binarySearch(arr []int, item int) (int, int) {
	low := 0
	high := len(arr) - 1
	count := 1
	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == item {
			return mid, count
		} else if arr[mid] > item {
			high = mid - 1
			count++
		} else {
			low = mid + 1
			count++
		}
	}
	return -1, 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	nStr, _ := reader.ReadString('\n')
	nInt, _ := strconv.Atoi(strings.TrimSpace(nStr))

	line, _ := reader.ReadString('\n')

	item, _ := reader.ReadString('\n')
	iInt, _ := strconv.Atoi(strings.TrimSpace(item))
	parts := strings.Fields(line)

	arr := make([]int, nInt)
	for i := 0; i < nInt; i++ {
		arr[i], _ = strconv.Atoi(parts[i])
	}

	fmt.Println(binarySearch(arr, iInt))
}
