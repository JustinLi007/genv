package main

import (
	"os"

	"github.com/JustinLi007/genv/internal/assert"
	"github.com/JustinLi007/genv/internal/commander"
	"github.com/JustinLi007/genv/internal/configs"
)

func main() {
	args := os.Args

	c, err := configs.NewConfigs()
	assert.NoErr(err, "failed to create configs")

	commander, err := commander.New(c)
	assert.NoErr(err, "failed to create commander")

	commander.Parse(args[1:])
}
