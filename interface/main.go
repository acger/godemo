package main

type Foo interface {
	Bar()
}

type F1 struct {
	Title string
}

func (f F1) Bar() {
	println(f.Title)
}

type F2 = F1

func main() {
	f1 := F1{
		Title: "f1",
	}

	f1.Bar()

	f2 := F2{
		Title: "f2",
	}

	f2.Bar()
}
