package main

type T2 struct {
	a int
}

func (t T2) Method1() {
}

func (t *T2) Method2() {
	t.a = 11
}

func main() {
	var t T2
	t.Method1() // ok
	t.Method2() // <=> (&t).M2()
	var pt = &T2{}
	pt.Method1() // <=> (*pt).M1()
	pt.Method2() // ok
}
