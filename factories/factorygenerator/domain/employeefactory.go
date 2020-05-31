package domain

// What if we want factories for specific roles?
// Approach 1: Functional
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{
			Name:         name,
			Position:     position,
			AnnualIncome: annualIncome,
		}
	}
}

// Approach 2: Structural
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{Position: position, AnnualIncome: annualIncome}
}

func (ef *EmployeeFactory) Create(name string) *Employee {
	return &Employee{
		Name:         name,
		Position:     ef.Position,
		AnnualIncome: ef.AnnualIncome,
	}
}
