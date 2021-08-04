package sqrtx_69

import (
	"fmt"
)

// (xn + a/xn) / 2
// 假定计算根号2，即a是2，x0默认为1
// (1+2/1)/2 = 1.5
// (1.5+2/1.5)/2=1.416666667
// (1.416666667+2/1.416666667)/2=1.414215686

func mySqrt(x int) int {
	if x == 1 {
		return x
	}
	min := 0
	max := x
	m := (max + min) / 2
	for (max-min) > 1 {
		if x/m < m {
			max = m
		} else {
			min = m
		}
		m = (max + min) / 2
	}
	return min
}

func Run() {
	fmt.Println(mySqrt(1))
	fmt.Println(mySqrt(12))
	fmt.Println(mySqrt(22))
	fmt.Println(mySqrt(32))
	fmt.Println(mySqrt(42))
}
