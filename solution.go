package main

import "fmt"

type SolutionImpl struct {
}

func (s SolutionImpl) Resolve() int {
	return s.answer1()
}

func (SolutionImpl) answer1() int {
	fmt.Println("hello world")
	return 0
}
