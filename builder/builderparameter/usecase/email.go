package usecase

import (
	"fmt"

	"github.com/isfanazha/go-design-patterns/builder/builderparameter/domain"
)

func SendEmail(action domain.Build) {
	builder := domain.EmailBuilder{}
	action(&builder)
	sendEmail(&builder.Email)
}

func sendEmail(email *domain.Email) {
	fmt.Println(fmt.Sprintf("Sending email from %v, to %v", email.From, email.To))
	fmt.Println(fmt.Sprintf("With subject: %v, and body: %v", email.Subject, email.Body))
}
