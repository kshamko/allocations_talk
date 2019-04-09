package example2

type x struct {
	data *string
}

func indirect() {
	str := "xxx"
	o := new(x)
	o.data = &str //BAD
}

func direct() {
	str := "xxx"
	o := &x{
		data: &str, //OK
	}
	_ = o
}
