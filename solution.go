package main

import "fmt"

type SolutionImpl struct {
}

func (s SolutionImpl) Resolve() int {
	answer := s.answer1([]int{1, 2})
	fmt.Printf("%v\n", answer)
	return 0
}

func (SolutionImpl) answer1(nums []int) []int {
	return nums
}
