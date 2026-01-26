package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/JustinLi007/genv/internal/action"
	"github.com/JustinLi007/genv/internal/locator"
	"github.com/JustinLi007/genv/internal/services"
	"github.com/JustinLi007/genv/internal/utils"
)

type HandlerProjects interface {
	Help(ar *action.ActionRequest)
	NewProject(ar *action.ActionRequest)
	GetProject(ar *action.ActionRequest)
	GetProjects(ar *action.ActionRequest)
	DeleteProject(ar *action.ActionRequest)
}

type handlerProjects struct {
	serviceProjects services.ServiceProjects
	locator         *locator.Locator
}

func NewHandlerProjects(serviceProjects services.ServiceProjects, locator *locator.Locator) HandlerProjects {
	hp := &handlerProjects{
		serviceProjects: serviceProjects,
		locator:         locator,
	}
	return hp
}

func (hp *handlerProjects) Help(ar *action.ActionRequest) {
	hp.locator.GetPrinter().Yellow([]byte(hp.serviceProjects.GetHelpMessage())).Println()
}

func (hp *handlerProjects) NewProject(ar *action.ActionRequest) {
	if ar == nil {
		return
	}

	targetDirs, ok := ar.Get("directory").(string)
	if !ok {
		return
	}

	req := make([]string, 0)
	for v := range strings.SplitSeq(targetDirs, ",") {
		if _, ap, err := utils.IsDirectory(v, true, true); err == nil {
			req = append(req, ap)
		}
	}

	err := hp.serviceProjects.InsertProject(req...)
	if err != nil {
		hp.locator.GetLogger().Error(err.Error())
	}
}

func (hp *handlerProjects) GetProject(ar *action.ActionRequest) {
	if ar == nil {
		return
	}

	targetDir, ok := ar.Get("directory").(string)
	if !ok {
		return
	}

	td := targetDir
	if num, err := strconv.ParseInt(targetDir, 10, 64); err == nil {
		td, err = hp.serviceProjects.SelectProjectById(num)
		if err != nil {
			return
		}
	}

	_, ap, err := utils.IsDirectory(td, true, true)
	if err != nil {
		return
	}
	td = ap

	if edit, ok := ar.Get("edit").(bool); ok && edit {
		ar.Set("directory", td)
		action.Redirect(ar, "tmux/directory")
		return
	}
	hp.locator.GetPrinter().Write([]byte(td)).Println()
}

func (hp *handlerProjects) GetProjects(ar *action.ActionRequest) {
	if ar == nil {
		return
	}

	p := ""
	s := ""
	c := ""
	if val, ok := ar.Get("prefix").(string); ok {
		p = val
	}
	if val, ok := ar.Get("suffix").(string); ok {
		s = val
	}
	if val, ok := ar.Get("contains").(string); ok {
		c = val
	}

	res, err := hp.serviceProjects.SelectProjects(p, s, c)
	if err != nil {
		return
	}

	prt := hp.locator.GetPrinter()
	for k, v := range res {
		prt.Write(fmt.Appendf(make([]byte, 0), "%d\t%s\t%s\n", k, "->", v))
	}
	prt.Print()
}

func (hp *handlerProjects) DeleteProject(ar *action.ActionRequest) {
	if ar == nil {
		return
	}

	targetDirs, ok := ar.Get("directory").(string)
	if !ok {
		return
	}

	req := make([]string, 0)
	for v := range strings.SplitSeq(targetDirs, ",") {
		if _, ap, err := utils.IsDirectory(v, true, true); err == nil {
			req = append(req, ap)
		}
	}

	err := hp.serviceProjects.DeleteProject(req...)
	hp.locator.GetLogger().Error(err.Error())
}
