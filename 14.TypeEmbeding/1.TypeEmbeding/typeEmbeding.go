package main


type animal interface {
	breath()
	walk()
}


type human interface {
	animal
	speak()
}

type base struct {
	num int
}

type container struct {
	human
	base
	str string
}


