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
		data: &str, //data var is initialized with the address of str var
	}
	_ = o
}
