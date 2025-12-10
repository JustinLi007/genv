package handlers

import (
	"github.com/JustinLi007/genv/internal/action"
	"github.com/JustinLi007/genv/internal/locator"
	"github.com/JustinLi007/genv/internal/services"
)

type HandlerGenv interface {
	Help(ar *action.ActionRequest)
}

type handlerGenv struct {
	serviceGenv services.ServiceGenv
	locator     *locator.Locator
}

func NewHandlerGenv(serviceGenv services.ServiceGenv, locator *locator.Locator) HandlerGenv {
	hg := &handlerGenv{
		serviceGenv: serviceGenv,
		locator:     locator,
	}
	return hg
}

func (hg *handlerGenv) Help(ar *action.ActionRequest) {
	hg.locator.GetPrinter().Yellow([]byte(hg.serviceGenv.GetHelpMessage())).Println()
}
