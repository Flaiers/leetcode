package main

import (
	"leetcode/daily-temperatures/stack"
)

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	dailyTemperatures(temperatures)
}

func dailyTemperatures(temperatures []int) []int {
	type idxTemp struct {
		idx  int
		temp int
	}

	days, s := make([]int, len(temperatures)), stack.New[idxTemp]()
	for i := len(temperatures) - 1; i >= 0; i-- {
		for s.Len() != 0 && s.Peek().temp <= temperatures[i] {
			s.Pop()
		}

		if s.Len() != 0 {
			days[i] = s.Peek().idx - i
		}
		s.Push(&idxTemp{
			idx:  i,
			temp: temperatures[i],
		})
	}

	return days
}
