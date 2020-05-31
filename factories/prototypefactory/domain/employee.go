package domain

type Employee struct {
	Name         string
	Position     string
	AnnualIncome int
}

const (
	Developer = iota
	Manager
)

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "Developer", 60000}
	case Manager:
		return &Employee{"", "Manager", 80000}
	default:
		panic("unsupported role")
	}
}
