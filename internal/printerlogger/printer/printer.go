package printer

import (
	"fmt"
	"strings"
	"text/tabwriter"
)

type PrinterComponent interface {
	Write(b []byte) PrinterComponent
	Blue(b []byte) PrinterComponent
	Yellow(b []byte) PrinterComponent
	Red(b []byte) PrinterComponent
	Print()
	Println()
	String() string
}

const (
	red        string = "\033[0;31m"
	yellow     string = "\033[0;33m"
	blue       string = "\033[0;34m"
	resetColor string = "\033[0m"
)

type printer struct {
	strBuilder *strings.Builder
	twriter    *tabwriter.Writer
}

func New() PrinterComponent {
	pt := &printer{
		strBuilder: &strings.Builder{},
	}
	pt.twriter = tabwriter.NewWriter(pt.strBuilder, 0, 0, 2, ' ', 0)
	return pt
}

func (pt *printer) Write(b []byte) PrinterComponent {
	fmt.Fprintf(pt.twriter, "%s", b)
	return pt
}

func (pt *printer) Blue(b []byte) PrinterComponent {
	fmt.Fprintf(pt.twriter, "%s%s%s", blue, b, resetColor)
	return pt
}

func (pt *printer) Yellow(b []byte) PrinterComponent {
	fmt.Fprintf(pt.twriter, "%s%s%s", yellow, b, resetColor)
	return pt
}

func (pt *printer) Red(b []byte) PrinterComponent {
	fmt.Fprintf(pt.twriter, "%s%s%s", red, b, resetColor)
	return pt
}

func (pt *printer) Print() {
	fmt.Print(pt)
}

func (pt *printer) Println() {
	fmt.Println(pt)
}

func (pt *printer) String() string {
	if err := pt.twriter.Flush(); err != nil {
		fmt.Printf("flush error: %s\n", err)
	}
	result := pt.strBuilder.String()
	pt.strBuilder.Reset()
	return result
}

type nullPrinter struct {
}

func NewNull() PrinterComponent {
	return &nullPrinter{}
}

func (np *nullPrinter) Write(b []byte) PrinterComponent {
	return np
}

func (np *nullPrinter) Blue(b []byte) PrinterComponent {
	return np
}

func (np *nullPrinter) Yellow(b []byte) PrinterComponent {
	return np
}

func (np *nullPrinter) Red(b []byte) PrinterComponent {
	return np
}

func (np *nullPrinter) Print() {
}

func (np *nullPrinter) Println() {
}

func (np *nullPrinter) String() string {
	return ""
}
