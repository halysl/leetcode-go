package guessnum_lcp01

import "fmt"

func game(guess []int, answer []int) int {
	count := 0
	for index, val := range guess {
		if answer[index] == val {
			count++
		}
	}
return count
}

func Run() {
	fmt.Println(game([]int{1,2,3}, []int{1,2,3}))
}