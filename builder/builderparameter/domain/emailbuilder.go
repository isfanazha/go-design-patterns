package domain

import "strings"

type EmailBuilder struct {
	Email Email
}

func (eb *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("email should contain @")
	}

	eb.Email.From = from
	return eb
}

func (eb *EmailBuilder) To(to string) *EmailBuilder {
	eb.Email.To = to
	return eb
}

func (eb *EmailBuilder) Subject(subject string) *EmailBuilder {
	eb.Email.Subject = subject
	return eb
}

func (eb *EmailBuilder) Body(body string) *EmailBuilder {
	eb.Email.Body = body
	return eb
}

// Need to pass pointer, because we want to fill the property data when function is executed
type Build func(builder *EmailBuilder)

