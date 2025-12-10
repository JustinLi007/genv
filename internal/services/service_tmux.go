package services

import (
	"fmt"

	database "github.com/JustinLi007/genv/internal/database/gob"
	"github.com/JustinLi007/genv/internal/locator"
	"github.com/JustinLi007/genv/internal/structures"
)

type ServiceTmux interface {
	Service
	GetTmuxScript() (*structures.Script, error)
}

type serviceTmux struct {
	db      database.Database
	locator *locator.Locator
}

func NewServiceTmux(db database.Database, locator *locator.Locator) ServiceTmux {
	st := &serviceTmux{
		db:      db,
		locator: locator,
	}
	return st
}

func (st *serviceTmux) GetHelpMessage() string {
	p := st.locator.GetPrinter()
	b := make([]byte, 0)
	b = fmt.Appendf(b, "USAGE:\n")
	b = fmt.Appendf(b, "\tgenv tmux [-h|--help]\n")
	b = fmt.Appendf(b, "\tgenv tmux -d <dir>\n")
	b = fmt.Appendf(b, "\nDESCRIPTION:\n")
	b = fmt.Appendf(b, "\tCreate or attach to a tmux session using a specified directory as the cwd.\n")
	b = fmt.Appendf(b, "\n\tUsage of '.' and '..' is allowed. If a directory is added using\n")
	b = fmt.Appendf(b, "\tthe projects action, then the integer mapped to the directory can be used.\n")
	b = fmt.Appendf(b, "\nFLAGS:\n")
	b = fmt.Appendf(b, "\t-h, --help\t\tprint usage info\n")
	b = fmt.Appendf(b, "\t-d, --directory\t\tan absolute path to a directory for the session's cwd.\n")
	p.Write(b)
	return p.String()
}

func (st *serviceTmux) GetTmuxScript() (*structures.Script, error) {
	p, err := st.db.ReadTmuxScriptPath()
	if err != nil {
		return nil, err
	}
	return structures.NewScript(p)
}
