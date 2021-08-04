package sumof2num_1

import "fmt"

func solution(nums []int, target int) []int {
	for index1, value := range nums {
		for index2:=index1+1;index2<len(nums);index2++ {
			if target - value == nums[index2] {
				return []int{index1, index2}
			}
		}
	}
return []int{}
}

func Run() {
fmt.Println(solution([]int{2, 7, 11, 15}, 9))
fmt.Println(solution([]int{3, 2, 4}, 6))
fmt.Println(solution([]int{-1, -2, -3, -4, -5}, -8))
}
