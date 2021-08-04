package kidshavecandy_1431

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := candies[0]
	for _, val := range candies {
		if val > max {
			max = val
		}
	}
	res := make([]bool, len(candies))
	for index, val := range candies {
		if val + extraCandies >= max {
			res[index] = true
		}
	}
return res
}

func Run() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
}
