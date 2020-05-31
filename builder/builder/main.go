package main

import (
	"fmt"

	"github.com/isfanazha/go-design-patterns/builder/builder/domain"
)

func main() {
	// Builder
	b := domain.NewHtmlBuilder("ul")
	b.AddChild("li", "hello")
	b.AddChild("li", "world")
	fmt.Print(b.String())
	fmt.Println()

	// Fluent (Chain)
	bf := domain.NewHtmlBuilder("ul")
	bf.AddChildFluent("li", "hello").
		AddChildFluent("li", "bro")

	fmt.Print(bf.String())

	// if you want to build a simple HTML paragraph
	// using a strings.Builder
	//hello := "hello"
	//sb := strings.Builder{}
	//sb.WriteString("<p>")
	//sb.WriteString(hello)
	//sb.WriteString("</p>")
	//fmt.Printf("%s\n", sb.String())
}
