package example3

type X struct {
	c int
}

func (x *X) add(i int) {
	x.c += i
}

func indirectCall() {
	x := new(X)

	f := x.add // BAD
	f(5)

	x.add(5) //GOOD
}
