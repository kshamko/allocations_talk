package example3

func callClosure() {
	var y int // BAD: y escapes
	func(p *int, x int) {
		*p = x
	}(&y, 42)
}

func callClosureDefer() {
	x := 0 // BAD: x escapes
	defer func(p *int) {
		*p = 1
	}(&x)
}
