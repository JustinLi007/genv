package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func IsDirectory(rawPath string, allowCwd, allowParent bool) (relative, absolute string, err error) {
	rp := strings.TrimSpace(rawPath)
	if rp == "" {
		return "", "", fmt.Errorf("empty path")
	}

	p := filepath.Clean(rp)
	if p == "." || p == ".." {
		if !allowCwd {
			return "", "", fmt.Errorf("cwd not allowed")
		}

		cwd, err := os.Getwd()
		if err != nil {
			return "", "", err
		}

		if p == "." {
			p = cwd
		} else if allowParent && p == ".." {
			p = filepath.Dir(cwd)
		}
	}

	fi, err := os.Stat(p)
	if err != nil {
		return "", "", err
	}

	if !fi.IsDir() {
		return "", "", fmt.Errorf("not a directory")
	}

	ap, err := filepath.Abs(p)
	if err != nil {
		return "", "", err
	}

	return p, ap, nil
}

func IsFile(rawPath string) (relative, absolute string, err error) {
	rp := strings.TrimSpace(rawPath)
	if rp == "" {
		return "", "", fmt.Errorf("empty path")
	}

	p := filepath.Clean(rp)
	if p == "." {
		return "", "", fmt.Errorf("not a file")
	}

	fi, err := os.Stat(p)
	if err != nil {
		return "", "", err
	}

	if fi.IsDir() {
		return "", "", fmt.Errorf("not a file")
	}

	ap, err := filepath.Abs(p)
	if err != nil {
		return "", "", err
	}

	return p, ap, nil
}

func CreateDirIfNotExist(filename string, perm os.FileMode) (bool, error) {
	_, _, err := IsDirectory(filename, true, true)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return false, err
	} else if errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(filename, perm); err != nil {
			return false, err
		} else {
			return true, nil
		}
	}
	return false, nil
}
