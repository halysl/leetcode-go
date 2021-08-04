package dohomework_1450

import "fmt"

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	count := 0
	for index, val := range startTime {
		if queryTime <= endTime[index] && queryTime >= val {
			count++
		}
	}
return count
}

func Run() {
	// startTime = [1,2,3], endTime = [3,2,7], queryTime = 4
	fmt.Println(busyStudent([]int{1,2,3},[]int{3,2,7}, 4))
}