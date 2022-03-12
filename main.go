package main

type Solution interface {
	Resolve() int
}

func main() {
	var s SolutionImpl = SolutionImpl{}
	s.Resolve()
}
