package structures

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/JustinLi007/genv/internal/utils"
)

type Script struct {
	interpreter string
	filename    string
	args        []string
}

func NewScript(filename string, args ...string) (*Script, error) {
	_, ap, err := utils.IsFile(filename)
	if err != nil {
		return nil, err
	}

	line, err := readInterpreterLine(ap)
	if err != nil {
		return nil, err
	}

	if !allowedInterpreters.has(line) {
		return nil, fmt.Errorf("'%s' not allowed", line)
	}

	s := &Script{
		interpreter: allowedInterpreters.get(line),
		filename:    ap,
		args:        []string{ap},
	}
	s.AddArgs(args...)

	return s, nil
}

func (s *Script) AddArgs(args ...string) {
	if len(args) <= 0 {
		return
	}
	s.args = append(s.args, args...)
}

func (s *Script) Run(ctx context.Context) ([]byte, error) {
	cmd := exec.CommandContext(ctx, s.interpreter, s.args...)
	return cmd.CombinedOutput()
}

func (s *Script) String() string {
	return fmt.Sprintf("%s -> %v\n", s.interpreter, s.args)
}

func readInterpreterLine(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(line), nil
}

var allowedInterpreters *interpreters = newInterpreters()

type interpreters struct {
	allowed map[string]string
}

func newInterpreters() *interpreters {
	ip := &interpreters{
		allowed: make(map[string]string),
	}
	ip.set("#!/usr/bin/env bash", "bash")
	ip.set("#!/bin/bash", "bash")
	ip.set("#!/usr/bin/env zsh", "zsh")
	ip.set("#!/bin/zsh", "zsh")
	ip.set("#!/bin/sh", "sh")
	return ip
}

func (ip *interpreters) set(key, value string) {
	if ip.allowed == nil {
		return
	}
	ip.allowed[key] = value
}

func (ip *interpreters) has(key string) bool {
	if ip.allowed == nil {
		return false
	}
	_, ok := ip.allowed[key]
	return ok
}

func (ip *interpreters) get(key string) string {
	if ip.has(key) {
		return ip.allowed[key]
	}
	return ""
}
