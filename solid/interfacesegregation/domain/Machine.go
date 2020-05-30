package domain

import "fmt"

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultifunctionPrinter struct{}

func (m *MultifunctionPrinter) Print(d Document) {
	fmt.Println("MultifunctionPrinter executes print")
}

func (m *MultifunctionPrinter) Fax(d Document) {
	fmt.Println("MultifunctionPrinter execute fax")
}

func (m *MultifunctionPrinter) Scan(d Document) {
	fmt.Println("MultifunctionPrinter execute scan")
}

// OldFashionedPrinter implement Machine, but it doesn't support Fax() and Scan() method.
// So we need to add deprecated to those functions. This is not good solution, we need to implement
// interface segregation principle.
type OldFashionedPrinter struct{}

func (o *OldFashionedPrinter) Print(d Document) {
	fmt.Println("OldFashionedPrinter executes print")
}

// Deprecated: ...
func (o *OldFashionedPrinter) Fax(d Document) {
	panic("operation not supported")
}

// Deprecated: ...
func (o *OldFashionedPrinter) Scan(d Document) {
	panic("operation not supported")
}

// Approach 1: Split into several interfaces
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

// PrinterOnly implement Printer interface, because PrinterOnly can only print.
type PrinterOnly struct{}

func (p *PrinterOnly) Print(d Document) {
	fmt.Println("PrinterOnly executes print")
}

// Approach 2: Combine interfaces
// Photocopier implement Printer and Scanner interface, because Photocopier can print and scan.
type Photocopier struct{}

func (p *Photocopier) Print(d Document) {
	fmt.Println("Photocopier executes print")
}

func (p *Photocopier) Scan(d Document) {
	fmt.Println("Photocopier executes scan")
}

// Approach 3: Interface combination + decorator
// MultiFunctionDevice interface embeds Printer and Scanner interface
type MultiFunctionDevice interface {
	Printer
	Scanner
}

// MultiFunctionMachine
type MultiFunctionMachine struct {
	Printer Printer
	Scanner Scanner
}

func (m *MultiFunctionMachine) Print(d Document) {
	m.Printer.Print(d) // decorator
}

func (m *MultiFunctionMachine) Scan(d Document) {
	m.Scanner.Scan(d) // decorator
}
