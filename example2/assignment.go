package example2

type x struct {
	data *string
}

// BAD
func indirect() {
	str := "xxx"
	o := new(x)
	o.data = &str
}

// OK
func direct() {
	str := "xxx"
	o := &x{
		data: &str,
	}
	_ = o
}
