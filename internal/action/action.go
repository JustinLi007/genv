package action

import (
	"fmt"
	"strings"

	"github.com/JustinLi007/genv/internal/assert"
)

type Action struct {
	Handler ActionHandler
}

func Redirect(ar *ActionRequest, route string) {
	if ar == nil {
		return
	}
	ar.Set("action", route)
	ar.ex.Handler.Perform(ar)
}

func (a *Action) Parse(args []string) {
	offset := 0
	ar := NewActionRequest(a)
	offset += getMethod(args[offset:], ar)
	offset += getAction(args[offset:], ar)
	assert.NotNil(ar, "executor: nil options")
	a.Handler.Perform(ar)
}

func getMethod(args []string, ar *ActionRequest) int {
	assert.True(len(args) > 0, "executor: no method specified")
	assert.NotNil(ar, "executor: nil options")

	methodStr := strings.ToLower(strings.TrimSpace(args[0]))

	offset := 1
	switch methodStr {
	case "--new":
		ar.Set("method", "new")
	case "--get":
		ar.Set("method", "get")
	case "--delete":
		ar.Set("method", "delete")
	default:
		ar.Set("method", "get")
		offset--
	}

	return offset
}

func getAction(args []string, ar *ActionRequest) int {
	assert.True(len(args) > 0, "executor: no action specified")
	assert.NotNil(ar, "executor: nil options")

	actionStr := strings.ToLower(strings.TrimSpace(args[0]))

	offset := 1
	switch actionStr {
	case "tmux":
		ar.Set("action", "tmux")
		offset += getTmuxFlags(args[offset:], ar)
	case "projects":
		ar.Set("action", "projects")
		offset += getProjectFlags(args[offset:], ar)
	default:
		ar.Set("action", "help")
	}

	return offset
}

func getProjectFlags(args []string, ar *ActionRequest) int {
	assert.NotNil(ar, "executor: nil options")

	n := len(args)
	if n == 0 {
		return 0
	}

	i := 0
	for i < n {
		cur := args[i]
		switch cur {
		case "-h", "--help":
			ar.Set("help", true)
		case "-e", "--edit":
			ar.Set("edit", true)
		case "-d", "--directory":
			assert.True(i+1 < n, fmt.Sprintf("invalid value for %s", cur))
			i++
			ar.Set("directory", args[i])
		case "--prefix":
			assert.True(i+1 < n, fmt.Sprintf("invalid value for %s", cur))
			i++
			ar.Set("prefix", args[i])
		case "--suffix":
			assert.True(i+1 < n, fmt.Sprintf("invalid value for %s", cur))
			i++
			ar.Set("suffix", args[i])
		case "--contains":
			assert.True(i+1 < n, fmt.Sprintf("invalid value for %s", cur))
			i++
			ar.Set("contains", args[i])
		default:
			assert.True(false, "unknown flag")
		}
		i++
	}

	if ar.Has("help") {
		if val, ok := ar.Get("action").(string); ok {
			ar.Set("action", fmt.Sprintf("%s/%s", val, "help"))
		}
	} else if ar.Has("directory") {
		if val, ok := ar.Get("action").(string); ok {
			ar.Set("action", fmt.Sprintf("%s/%s", val, "directory"))
		}
	}

	return i
}

func getTmuxFlags(args []string, ar *ActionRequest) int {
	assert.NotNil(ar, "executor: nil options")

	n := len(args)
	if n == 0 {
		return 0
	}

	i := 0
	for i < n {
		cur := args[i]
		switch cur {
		case "-h", "--help":
			ar.Set("help", true)
		case "-d", "--directory":
			assert.True(i+1 < n, fmt.Sprintf("invalid value for %s", cur))
			i++
			ar.Set("directory", args[i])
		default:
			assert.True(false, "unknown flag")
		}
		i++
	}

	if ar.Has("help") {
		if val, ok := ar.Get("action").(string); ok {
			ar.Set("action", fmt.Sprintf("%s/%s", val, "help"))
		}
	} else if ar.Has("directory") {
		if val, ok := ar.Get("action").(string); ok {
			ar.Set("action", fmt.Sprintf("%s/%s", val, "directory"))
		}
	}

	return i
}
