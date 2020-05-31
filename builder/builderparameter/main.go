package main

import (
	"github.com/isfanazha/go-design-patterns/builder/builderparameter/domain"
	"github.com/isfanazha/go-design-patterns/builder/builderparameter/usecase"
)

func main() {
	usecase.SendEmail(func(b *domain.EmailBuilder) {
		b.From("foo@bar.com").
			To("bar@baz.com").
			Subject("Meeting").
			Body("Hello, do you want to meet?")
	})
}
