package main

func main() {

}

func makeFooGoodCase() (Foo, error) {
	f := Foo{
		"1",
		1,
	}
	return f, nil
}

func makeFooBadCase(f *Foo) error {
	f.f1 = "1"
	f.f2 = 1
	return nil
}

type Foo struct {
	f1 string
	f2 int
}
