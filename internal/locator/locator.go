package locator

import (
	"github.com/JustinLi007/genv/internal/assert"
	"github.com/JustinLi007/genv/internal/printerlogger/logger"
	"github.com/JustinLi007/genv/internal/printerlogger/printer"
)

type Locator struct {
	loggerComp  logger.LoggerComponent
	printerComp printer.PrinterComponent

	nullLoggerComp  logger.LoggerComponent
	nullPrinterComp printer.PrinterComponent
}

func New() *Locator {
	locator := &Locator{
		nullLoggerComp:  logger.NewNull(),
		nullPrinterComp: printer.NewNull(),
	}

	msg := "null service cannot be nil"
	assert.NotNil(locator.nullLoggerComp, msg)
	assert.NotNil(locator.nullPrinterComp, msg)

	return locator
}

func (l *Locator) RegisterLogger(comp logger.LoggerComponent) {
	if assert.IsNil(comp) {
		l.loggerComp = l.nullLoggerComp
		return
	}
	l.loggerComp = comp
}

func (l *Locator) GetLogger() logger.LoggerComponent {
	return l.loggerComp
}

func (l *Locator) RegisterPrinter(comp printer.PrinterComponent) {
	if assert.IsNil(comp) {
		l.printerComp = l.nullPrinterComp
		return
	}
	l.printerComp = comp
}

func (l *Locator) GetPrinter() printer.PrinterComponent {
	return l.printerComp
}
