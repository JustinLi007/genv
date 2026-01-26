package commander

import (
	"github.com/JustinLi007/genv/internal/action"
	"github.com/JustinLi007/genv/internal/assert"
	"github.com/JustinLi007/genv/internal/configs"
	database "github.com/JustinLi007/genv/internal/database/gob"
	"github.com/JustinLi007/genv/internal/handlers"
	"github.com/JustinLi007/genv/internal/locator"
	"github.com/JustinLi007/genv/internal/printerlogger/logger"
	"github.com/JustinLi007/genv/internal/printerlogger/printer"
	"github.com/JustinLi007/genv/internal/services"
)

type commander struct {
	handlerGenv     handlers.HandlerGenv
	handlerTmux     handlers.HandlerTmux
	handlerProjects handlers.HandlerProjects
}

func New(c *configs.Config) (*action.Action, error) {
	locator := locator.New()
	assert.NotNil(locator, "failed to create locator")

	locator.RegisterLogger(logger.New())
	locator.RegisterPrinter(printer.New())

	db, err := database.NewGobDatabase(
		c.GobDbFilename(),
		c.TmuxScriptFilename(),
		c.PermDir(),
		locator,
	)
	if err != nil {
		return nil, err
	}

	serviceGenv := services.NewServiceGenv(locator)
	serviceTmux := services.NewServiceTmux(db, locator)
	serviceProjects := services.NewServiceProjects(db, locator)

	handlerGenv := handlers.NewHandlerGenv(serviceGenv, locator)
	handlerTmux := handlers.NewHandlerTmux(serviceTmux, locator)
	handlerProjects := handlers.NewHandlerProjects(serviceProjects, locator)

	cmder := commander{
		handlerGenv:     handlerGenv,
		handlerTmux:     handlerTmux,
		handlerProjects: handlerProjects,
	}

	ex := &action.Action{
		Handler: cmder.RegisterActions(),
	}

	return ex, nil
}

func (c *commander) RegisterActions() *action.Mux {
	mux := action.NewMux()

	mux.Register("get", "help", c.handlerGenv.Help)

	mux.Register("get", "tmux/help", c.handlerTmux.Help)
	mux.Register("get", "tmux/directory", c.handlerTmux.CreateSession)

	mux.Register("new", "projects/directory", c.handlerProjects.NewProject)
	mux.Register("get", "projects", c.handlerProjects.GetProjects)
	mux.Register("get", "projects/directory", c.handlerProjects.GetProject)
	mux.Register("get", "projects/help", c.handlerProjects.Help)
	mux.Register("delete", "projects/directory", c.handlerProjects.DeleteProject)

	return mux
}
