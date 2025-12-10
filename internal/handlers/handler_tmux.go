package handlers

import (
	"context"

	"github.com/JustinLi007/genv/internal/action"
	"github.com/JustinLi007/genv/internal/locator"
	"github.com/JustinLi007/genv/internal/services"
	"github.com/JustinLi007/genv/internal/utils"
)

type HandlerTmux interface {
	Help(ar *action.ActionRequest)
	CreateSession(ar *action.ActionRequest)
}

type handlerTmux struct {
	serviceTmux services.ServiceTmux
	locator     *locator.Locator
}

func NewHandlerTmux(serviceTmux services.ServiceTmux, locator *locator.Locator) HandlerTmux {
	ht := &handlerTmux{
		serviceTmux: serviceTmux,
		locator:     locator,
	}
	return ht
}

func (ht *handlerTmux) Help(ar *action.ActionRequest) {
	ht.locator.GetPrinter().Yellow([]byte(ht.serviceTmux.GetHelpMessage())).Println()
}

func (ht *handlerTmux) CreateSession(ar *action.ActionRequest) {
	if ar == nil {
		return
	}

	targetDir, ok := ar.Get("directory").(string)
	if !ok {
		return
	}

	_, ap, err := utils.IsDirectory(targetDir, true, true)
	if err != nil {
		return
	}

	script, err := ht.serviceTmux.GetTmuxScript()
	if err != nil {
		ht.locator.GetLogger().Error(err.Error())
		return
	}
	script.AddArgs(ap)
	if output, err := script.Run(context.Background()); err != nil {
		ht.locator.GetLogger().Error(err.Error())
	} else {
		ht.locator.GetPrinter().Write(output).Println()
	}
}
