package sumof2num_167

import "fmt"

func solution(numbers []int, target int) []int {
	low, high := 0,len(numbers)-1
	for low < high {
		if numbers[low] + numbers[high] == target {
			return []int{low+1, high+1}
		} else if numbers[low] + numbers[high] > target {
			high--
		} else {
			low++
		}
	}
	return []int{}
}

func Run() {
	fmt.Println(solution([]int{2, 7, 11, 15}, 9))
	fmt.Println(solution([]int{2, 3, 4}, 6))
	fmt.Println(solution([]int{-5, -4, -3, -2, -1}, -8))
}
