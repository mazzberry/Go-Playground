package main



type taxSystem interface {
	calculateTax()
}

type indianTax struct {
	taxPercentage int
	income int
}

func (i *indianTax) calculateTax() int {
	tax := (i.income * i.taxPercentage) / 100
	return tax
}

type singaporeTax struct {
	taxPercentage int
	income int
}

func (s *singaporeTax) calculateTax() int {
	tax := (s.income * s.taxPercentage) / 100 
	return tax
}
