// Реализовать паттерн «адаптер» на любом примере.
package main

import (
	"fmt"
)

type LegacyPrinter interface {
	Print(s string) string
}

type MyPrinter struct{}

func (mp *MyPrinter) Print(s string) string {
	return fmt.Sprintf("Legacy Printer: %s", s)
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
}

func (pa *PrinterAdapter) Print(s string) string {
	return pa.OldPrinter.Print(s)
}

func main() {
	printer := &MyPrinter{}
	adapter := &PrinterAdapter{OldPrinter: printer}
	fmt.Println(adapter.Print("Hello, world!"))
}
