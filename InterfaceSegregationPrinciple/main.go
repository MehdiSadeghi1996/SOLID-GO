package main

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type OldFashionedPrinter struct {
	// ...
}

func (o OldFashionedPrinter) Print(d Document) {
	// اینجا اوکیه
}

func (o OldFashionedPrinter) Fax(d Document) {
	panic("نمیتواند پیاده سازی شود چون قدیمیه")
}

// Deprecated: ...
func (o OldFashionedPrinter) Scan(d Document) {
	panic("نمیتواند پیاده سازی شود چون قدیمیه")
}

// Printer کوچک کردن اینترفیس ها به اینترفیس های کوچک تر
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MultiFunctionDevice interface {
	Printer
	Scanner
}

type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

func main() {

}
