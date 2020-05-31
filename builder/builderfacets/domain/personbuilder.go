package domain

type PersonBuilder struct {
	Person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{
		Person: &Person{},
	}
}

func (pb *PersonBuilder) Build() *Person {
	return pb.Person
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (pb *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*pb}
}

func (pjb *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	pjb.Person.CompanyName = companyName
	return pjb
}

func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	pjb.Person.Position = position
	return pjb
}

func (pjb *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	pjb.Person.AnnualIncome = annualIncome
	return pjb
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (pb *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*pb}
}

func (pab *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	pab.Person.StreetAddress = streetAddress
	return pab
}

func (pab *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	pab.Person.City = city
	return pab
}

func (pab *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	pab.Person.Postcode = postcode
	return pab
}
