package services

import (
	"fmt"

	database "github.com/JustinLi007/genv/internal/database/gob"
	"github.com/JustinLi007/genv/internal/locator"
)

type ServiceProjects interface {
	Service
	InsertProject(items ...string) error
	SelectProjects(prefix, suffix, contains string) (map[int64]string, error)
	SelectProjectById(id int64) (string, error)
	DeleteProject(items ...string) error
}

type serviceProjects struct {
	db      database.Database
	locator *locator.Locator
}

func NewServiceProjects(db database.Database, locator *locator.Locator) ServiceProjects {
	sp := &serviceProjects{
		db:      db,
		locator: locator,
	}
	return sp
}

func (sp *serviceProjects) GetHelpMessage() string {
	p := sp.locator.GetPrinter()
	b := make([]byte, 0)
	b = fmt.Appendf(b, "USAGE:\n")
	b = fmt.Appendf(b, "\tgenv projects\n")
	b = fmt.Appendf(b, "\tgenv projects [-h|--help]\n")
	b = fmt.Appendf(b, "\tgenv projects [--prefix <string>|--suffix <string>|--contains <string>]\n")
	b = fmt.Appendf(b, "\tgenv projects [-d <dir>|--directory <dir>]\n")
	b = fmt.Appendf(b, "\nDESCRIPTION:\n")
	b = fmt.Appendf(b, "\tMark and unmark directories.\n")
	b = fmt.Appendf(b, "\nFLAGS:\n")
	b = fmt.Appendf(b, "\t-h, --help\t\tprint usage info\n")
	b = fmt.Appendf(b, "\t-d, --directory\t\tdirectory to mark or unmark\n")
	b = fmt.Appendf(b, "\t--prefix\t\tsearch marked directories that satisfy the prefix.\n")
	b = fmt.Appendf(b, "\t--suffix\t\tsearch marked directories that satisfy the suffix.\n")
	b = fmt.Appendf(b, "\t--contains\t\tsearch marked directories that contains a substring.\n")
	p.Write(b)
	return p.String()
}

func (sp *serviceProjects) InsertProject(items ...string) error {
	trie, err := sp.db.ReadProjectsData()
	if err != nil {
		return err
	}
	for _, v := range items {
		trie.Insert(v)
	}
	return sp.db.WriteProjectsData(trie)
}

func (sp *serviceProjects) SelectProjects(prefix, suffix, contains string) (map[int64]string, error) {
	trie, err := sp.db.ReadProjectsData()
	if err != nil {
		return nil, err
	}
	return trie.Search(prefix, suffix, contains)
}

func (sp *serviceProjects) SelectProjectById(id int64) (string, error) {
	trie, err := sp.db.ReadProjectsData()
	if err != nil {
		return "", err
	}
	return trie.GetByNum(id), nil
}

func (sp *serviceProjects) DeleteProject(items ...string) error {
	trie, err := sp.db.ReadProjectsData()
	if err != nil {
		return err
	}
	for _, v := range items {
		trie.Remove(v)
	}
	return sp.db.WriteProjectsData(trie)
}
