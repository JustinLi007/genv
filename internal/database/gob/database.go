package gob

import (
	"bufio"
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/JustinLi007/genv/internal/locator"
	"github.com/JustinLi007/genv/internal/utils"
)

type Database interface {
	ReadProjectsData() (*ProjectsTrie, error)
	WriteProjectsData(projectsData *ProjectsTrie) error
	ReadTmuxScriptPath() (string, error)
}

type gobDatabase struct {
	dbFilename         string
	tmuxScriptFilename string
	permDir            os.FileMode
	mtx                *sync.RWMutex
	locator            *locator.Locator
}

func NewGobDatabase(filename, tmuxScriptFilename string, permDir os.FileMode, locator *locator.Locator) (Database, error) {
	newDatabase := &gobDatabase{
		dbFilename:         filename,
		tmuxScriptFilename: tmuxScriptFilename,
		permDir:            permDir,
		mtx:                &sync.RWMutex{},
		locator:            locator,
	}
	if err := newDatabase.ensureGobDatabase(); err != nil {
		return nil, err
	}
	return newDatabase, nil
}

func (d *gobDatabase) ReadProjectsData() (*ProjectsTrie, error) {
	f, err := os.Open(d.dbFilename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var projectsTrie ProjectsTrie
	dec := gob.NewDecoder(f)
	err = dec.Decode(&projectsTrie)
	if err != nil {
		return nil, err
	}

	return &projectsTrie, nil
}

func (d *gobDatabase) WriteProjectsData(projectsData *ProjectsTrie) error {
	f, err := os.Create(d.dbFilename)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := bufio.NewWriter(f)
	defer buf.Flush()

	enc := gob.NewEncoder(buf)
	return enc.Encode(projectsData)
}

func (d *gobDatabase) ReadTmuxScriptPath() (string, error) {
	return d.tmuxScriptFilename, nil
}

func (d *gobDatabase) ensureGobDatabase() error {
	f, err := os.Open(d.dbFilename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return d.createDatabase()
		}
		return err
	}
	defer f.Close()
	return nil
}

func (d *gobDatabase) createDatabase() error {
	if created, err := utils.CreateDirIfNotExist(d.dbFilename, d.permDir); err != nil {
		return err
	} else if created {
		p := d.locator.GetPrinter()
		p.Write(fmt.Appendf(make([]byte, 0), "created %s\n", d.dbFilename))
		p.Print()
	}

	projectsTrie := NewProjectsTrie()
	return d.WriteProjectsData(projectsTrie)
}
