package configs

import (
	"os"
	"path/filepath"
)

type Config struct {
	progname           string
	genvDir            string
	homeDir            string
	gobDbFilename      string
	tmuxScriptFilename string
	permDir            os.FileMode
	permFile           os.FileMode
}

func NewConfigs() (*Config, error) {
	c := &Config{
		progname:           "genv",
		genvDir:            ".genv",
		homeDir:            "",
		gobDbFilename:      "",
		tmuxScriptFilename: "",
		permDir:            0o775,
		permFile:           0o666,
	}

	userHome, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	c.homeDir = userHome
	c.gobDbFilename = filepath.Join(userHome, c.genvDir, ".data/db")
	c.tmuxScriptFilename = filepath.Join(userHome, c.genvDir, "action/tmux.sh")

	return c, nil
}

func (c *Config) Progname() string {
	return c.progname
}

func (c *Config) GenvDir() string {
	return c.genvDir
}

func (c *Config) HomeDir() string {
	return c.homeDir
}

func (c *Config) GobDbFilename() string {
	return c.gobDbFilename
}

func (c *Config) TmuxScriptFilename() string {
	return c.tmuxScriptFilename
}

func (c *Config) PermDir() os.FileMode {
	return c.permDir
}

func (c *Config) PermFile() os.FileMode {
	return c.permFile
}
