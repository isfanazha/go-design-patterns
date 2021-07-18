package main

import (
	"fmt"
	"strings"

	"github.com/isfanazha/go-design-patterns/solid-principles/single-responsibility/domain"
	"github.com/isfanazha/go-design-patterns/solid-principles/single-responsibility/persistance"
	"github.com/isfanazha/go-design-patterns/solid-principles/single-responsibility/util"
)

func main() {
	// Initialize Journal struct.
	// Read the note in domain/journal.go
	j := domain.Journal{}
	j.AddEntry("Journal A")
	j.AddEntry("Journal B")

	fmt.Println(strings.Join(j.Entries, "\n"))

	// This is not using single responsibility principle
	// j.SaveToFile("journal.txt")

	// Approach 1: Using separate function
	util.SaveToFile(&j, "journal.txt")

	// Approach 2: Using separate type
	p := persistance.Persistence{"\n"}
	p.SaveToFile(&j, "journal.txt")
}
