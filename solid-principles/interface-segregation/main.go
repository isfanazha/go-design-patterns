package main

import (
	"fmt"

	"github.com/isfanazha/go-design-patterns/solid-principles/interface-segregation/domain"
)

func main() {
	doc := domain.Document{}

	// MultifunctionPrinter
	mfp := domain.MultifunctionPrinter{}
	mfp.Print(doc)
	mfp.Scan(doc)
	mfp.Print(doc)

	// OldFashionedPrinter
	ofp := domain.OldFashionedPrinter{}
	fmt.Println()
	ofp.Print(doc)
	// Will be panic
	// ofp.Scan(doc)
	// ofp.Fax(doc)

	// PrinterOnly
	p := domain.PrinterOnly{}
	fmt.Println()
	p.Print(doc)

	// Photocopier
	pco := domain.Photocopier{}
	fmt.Println()
	pco.Print(doc)
	pco.Scan(doc)

	// MultiFunctionMachine
	mfm := domain.MultiFunctionMachine{}
	mfm.Printer = &p // using PrinterOnly machine, or you can use Photocopier
	mfm.Scanner = &pco // using  Photocopier machine

	fmt.Println()
	mfm.Print(doc) 
	mfm.Scan(doc)
}
