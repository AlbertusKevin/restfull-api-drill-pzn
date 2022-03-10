package zsimple

type Foo struct {
}

func NewFoo() *Foo {
	return &Foo{}
}

type Bar struct {
}

func NewBar() *Bar {
	return &Bar{}
}

// misal, ada sebuah struct yang fieldnya butuh struct Foo dan Bar
type FooBar struct {
	*Bar
	*Foo
}

// biasanya, kita akan membuat sebuah function provider, yang nantinya kita inject Foo dan Barnya lewat function provider ini
func NewFooBar(bar *Bar, foo *Foo) *FooBar {
	return &FooBar{Bar: bar, Foo: foo}
}

// namun kita juga bisa melakukannya langsung menggunakan struct provider ketika proses Initialized
