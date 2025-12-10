package services

import (
	"fmt"

	"github.com/JustinLi007/genv/internal/locator"
)

type ServiceGenv interface {
	Service
}

type serviceGenv struct {
	locator *locator.Locator
}

func NewServiceGenv(locator *locator.Locator) ServiceGenv {
	sg := &serviceGenv{
		locator: locator,
	}
	return sg
}

func (sg *serviceGenv) GetHelpMessage() string {
	p := sg.locator.GetPrinter()
	b := make([]byte, 0)
	b = fmt.Appendf(b, "USAGE:\n")
	b = fmt.Appendf(b, "\tgenv <ACTION> [-h|--help]\n")
	b = fmt.Appendf(b, "\nDESCRIPTION:\n")
	b = fmt.Appendf(b, "\tDev env convenience tool.\n")
	b = fmt.Appendf(b, "\nACTIONS:\n")
	b = fmt.Appendf(b, "\ttmux\t\tcreate tmux session.\n")
	b = fmt.Appendf(b, "\tprojects\t\tmanage project directories.\n")
	b = fmt.Appendf(b, "\nFLAGS:\n")
	b = fmt.Appendf(b, "\t-h, --help\t\tprint usage info.\n")
	p.Write(b)
	return p.String()
}
