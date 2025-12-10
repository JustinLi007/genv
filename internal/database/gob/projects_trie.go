package gob

import (
	"errors"
	"fmt"
	"strings"

	"github.com/JustinLi007/genv/internal/structures"
)

var ErrNoResult = errors.New("no results")

const trieDelimiter = "{"

type ProjectsTrie struct {
	Counter  int64
	Root     *structures.TrieNode
	CountMap map[int64]string
}

func NewProjectsTrie() *ProjectsTrie {
	pt := &ProjectsTrie{
		Counter:  0,
		Root:     structures.NewTrieNode(""),
		CountMap: make(map[int64]string),
	}
	return pt
}

func (pt *ProjectsTrie) prepareValue(value string) []string {
	val := strings.TrimSpace(value)
	if val == "" {
		return make([]string, 0)
	}

	_, err := pt.Search(val, val, "")
	if err == nil {
		return make([]string, 0)
	}

	n := len(val)
	result := make([][]string, n+1)
	for i := n - 1; i >= 0; i-- {
		result[i] = append(result[i], string(val[i]))
		for j := range len(result[i+1]) {
			result[i] = append(result[i], fmt.Sprintf("%s%s", string(val[i]), result[i+1][j]))
		}
	}

	final := make([]string, 0)
	for _, res := range result {
		for _, v := range res {
			final = append(final, fmt.Sprintf("%s%s%s", v, trieDelimiter, val))
		}
	}
	final = append(final, fmt.Sprintf("%s%s", trieDelimiter, val))
	return final
}

func (pt *ProjectsTrie) GetCount() int64 {
	defer func() {
		pt.Counter++
	}()
	return pt.Counter
}

func (pt *ProjectsTrie) GetByNum(n int64) string {
	if val, ok := pt.CountMap[n]; ok {
		return val
	}
	return ""
}

func (pt *ProjectsTrie) Insert(value string) {
	toInsert := pt.prepareValue(value)
	if len(toInsert) <= 0 {
		fmt.Println("already exists")
		return
	}

	count := pt.GetCount()
	for _, v := range toInsert {
		pt.insert(v, count)
		if s, ok := strings.CutPrefix(v, trieDelimiter); ok {
			pt.CountMap[count] = s
		}
	}
}

func (pt *ProjectsTrie) insert(value string, count int64) {
	if pt.Root == nil {
		return
	}

	val := strings.TrimSpace(value)
	if val == "" {
		return
	}

	cur := pt.Root
	for v := range strings.SplitSeq(val, "") {
		if cur.Has(v) {
			cur = cur.Get(v)
			continue
		}
		node := structures.NewTrieNode(v)
		cur.Set(v, node, false)
		cur = cur.Get(v)
	}
	cur.IsEnd = true
	cur.Num = count
}

func (pt *ProjectsTrie) Remove(value string) {
	toRemove := pt.prepareValue(value)
	for _, v := range toRemove {
		pt.remove(v)
	}
}

func (pt *ProjectsTrie) remove(value string) {
	if pt.Root == nil {
		return
	}

	val := strings.TrimSpace(value)
	if val == "" {
		return
	}

	stack := make([]*structures.TrieNode, 0)

	cur := pt.Root
	for v := range strings.SplitSeq(val, "") {
		if !cur.Has(v) {
			return
		}
		stack = append(stack, cur)
		cur = cur.Get(v)
	}

	check := false
	for len(stack) > 0 && cur.Children == nil {
		// don't delete even if children is nil when popping up to a node marked IsEnd
		if check && cur.IsEnd {
			break
		}
		prev := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		prev.Delete(cur.Value)
		cur = prev
		check = true
	}
}

func (pt *ProjectsTrie) Search(prefix, suffix, contains string) (map[int64]string, error) {
	if pt.Root == nil {
		return nil, fmt.Errorf("nil root")
	}

	num := 0
	if contains != "" {
		num += 1
	}
	if suffix != "" {
		num += 2
	}
	if prefix != "" {
		num += 4
	}

	switch num {
	case 0: // none
		cur := pt.Root.Get(trieDelimiter)
		result := make(map[int64]string)
		display(cur, result, make([]string, 0))
		for k := range result {
			if val, ok := result[k]; ok {
				result[k] = strings.TrimPrefix(val, trieDelimiter)
			}
		}
		if len(result) == 0 {
			return nil, ErrNoResult
		}
		return result, nil
	case 1: // contains
		contains = strings.TrimSpace(contains)
		cur, _, err := pt.search(pt.Root, contains)
		if err != nil {
			return nil, err
		}
		if !cur.Has(trieDelimiter) {
			return nil, err
		}

		cur = cur.Get(trieDelimiter)
		result := make(map[int64]string)
		display(cur, result, make([]string, 0))
		for k := range result {
			if val, ok := result[k]; ok {
				result[k] = strings.TrimPrefix(val, trieDelimiter)
			}
		}
		if len(result) == 0 {
			return nil, ErrNoResult
		}
		return result, nil
	case 2: // suffix
		suffix = strings.TrimSpace(suffix)
		cur, _, err := pt.search(pt.Root, suffix)
		if err != nil {
			return nil, err
		}
		if !cur.Has(trieDelimiter) {
			return nil, err
		}

		discard := make([]int64, 0)
		cur = cur.Get(trieDelimiter)
		result := make(map[int64]string)
		display(cur, result, make([]string, 0))
		for k := range result {
			if val, ok := result[k]; ok {
				s := strings.TrimPrefix(val, trieDelimiter)
				result[k] = s
				if !strings.HasSuffix(s, suffix) {
					discard = append(discard, k)
				}
			}
		}
		for _, v := range discard {
			delete(result, v)
		}
		if len(result) == 0 {
			return nil, ErrNoResult
		}
		return result, nil
	case 3: // suffix and contains
		suffix = strings.TrimSpace(suffix)
		contains = strings.TrimSpace(contains)
		cur, _, err := pt.search(pt.Root, contains)
		if err != nil {
			return nil, err
		}
		if !cur.Has(trieDelimiter) {
			return nil, err
		}

		discard := make([]int64, 0)
		cur = cur.Get(trieDelimiter)
		result := make(map[int64]string)
		display(cur, result, make([]string, 0))
		for k := range result {
			if val, ok := result[k]; ok {
				s := strings.TrimPrefix(val, trieDelimiter)
				result[k] = s
				if !strings.HasSuffix(s, suffix) {
					discard = append(discard, k)
				}
			}
		}
		for _, v := range discard {
			delete(result, v)
		}
		if len(result) == 0 {
			return nil, ErrNoResult
		}
		return result, nil
	case 4: // prefix
		prefix = strings.TrimSpace(prefix)
		cur, _, err := pt.search(pt.Root, prefix)
		if err != nil {
			return nil, err
		}
		if !cur.Has(trieDelimiter) {
			return nil, err
		}

		discard := make([]int64, 0)
		cur = cur.Get(trieDelimiter)
		result := make(map[int64]string)
		display(cur, result, make([]string, 0))
		for k := range result {
			if val, ok := result[k]; ok {
				s := strings.TrimPrefix(val, trieDelimiter)
				result[k] = s
				if !strings.HasPrefix(s, prefix) {
					discard = append(discard, k)
				}
			}
		}
		for _, v := range discard {
			delete(result, v)
		}
		if len(result) == 0 {
			return nil, ErrNoResult
		}
		return result, nil
	case 5: // prefix and contains
		prefix = strings.TrimSpace(prefix)
		contains = strings.TrimSpace(contains)
		cur, _, err := pt.search(pt.Root, contains)
		if err != nil {
			return nil, err
		}
		if !cur.Has(trieDelimiter) {
			return nil, err
		}

		cur = cur.Get(trieDelimiter)
		cur, partial, err := pt.search(cur, prefix)
		if err != nil {
			return nil, err
		}
		partial = partial[1 : len(partial)-1]
		partialStr := strings.Join(partial, "")

		result := make(map[int64]string)
		display(cur, result, make([]string, 0))
		for k := range result {
			if val, ok := result[k]; ok {
				result[k] = fmt.Sprintf("%s%s", partialStr, val)
			}
		}
		if len(result) == 0 {
			return nil, ErrNoResult
		}
		return result, nil
	case 6: // prefix and suffix
		prefix = strings.TrimSpace(prefix)
		suffix = strings.TrimSpace(suffix)
		cur, _, err := pt.search(pt.Root, suffix)
		if err != nil {
			return nil, err
		}
		if !cur.Has(trieDelimiter) {
			return nil, err
		}

		cur = cur.Get(trieDelimiter)
		cur, partial, err := pt.search(cur, prefix)
		if err != nil {
			return nil, err
		}
		partial = partial[1 : len(partial)-1]
		partialStr := strings.Join(partial, "")

		discard := make([]int64, 0)
		result := make(map[int64]string)
		display(cur, result, make([]string, 0))
		for k := range result {
			if val, ok := result[k]; ok {
				s := fmt.Sprintf("%s%s", partialStr, val)
				result[k] = fmt.Sprintf("%s%s", partialStr, val)
				if !strings.HasSuffix(s, suffix) {
					discard = append(discard, k)
				}
			}
		}
		for _, v := range discard {
			delete(result, v)
		}
		if len(result) == 0 {
			return nil, ErrNoResult
		}
		return result, nil
	case 7: // all
		prefix = strings.TrimSpace(prefix)
		suffix = strings.TrimSpace(suffix)
		contains = strings.TrimSpace(contains)
		cur, _, err := pt.search(pt.Root, contains)
		if err != nil {
			return nil, err
		}
		if !cur.Has(trieDelimiter) {
			return nil, err
		}

		cur = cur.Get(trieDelimiter)
		cur, partial, err := pt.search(cur, prefix)
		if err != nil {
			return nil, err
		}
		partial = partial[1 : len(partial)-1]
		partialStr := strings.Join(partial, "")

		discard := make([]int64, 0)
		result := make(map[int64]string)
		display(cur, result, make([]string, 0))
		for k := range result {
			if val, ok := result[k]; ok {
				s := fmt.Sprintf("%s%s", partialStr, val)
				result[k] = s
				if !strings.HasSuffix(s, suffix) {
					discard = append(discard, k)
				}
			}
		}
		for _, v := range discard {
			delete(result, v)
		}
		if len(result) == 0 {
			return nil, ErrNoResult
		}
		return result, nil
	default:
		return nil, fmt.Errorf("invalid search state")
	}
}

func (pt *ProjectsTrie) search(root *structures.TrieNode, value string) (*structures.TrieNode, []string, error) {
	if root == nil {
		return nil, nil, fmt.Errorf("nil root")
	}

	val := strings.TrimSpace(value)
	if val == "" {
		return nil, nil, fmt.Errorf("empty value")
	}

	pathTaken := make([]string, 0)
	cur := root
	for v := range strings.SplitSeq(val, "") {
		if !cur.Has(v) {
			return nil, nil, fmt.Errorf("does not exist: '%s'", v)
		}
		pathTaken = append(pathTaken, cur.Value)
		cur = cur.Get(v)
	}
	pathTaken = append(pathTaken, cur.Value)

	return cur, pathTaken, nil
}

func (pt *ProjectsTrie) Display() {
	if pt.Root == nil {
		return
	}

	result := make(map[int64]string)
	temp := make([]string, 0)
	display(pt.Root.Get(trieDelimiter), result, temp)
	for k, v := range result {
		fmt.Printf("%-5d -> %s\n", k, strings.TrimPrefix(v, trieDelimiter))
	}
}

func display(root *structures.TrieNode, result map[int64]string, temp []string) {
	if root == nil {
		return
	}
	temp = append(temp, root.Value)

	if root.IsEnd {
		result[root.Num] = strings.Join(temp, "")
	}

	for _, v := range root.Children {
		display(v, result, temp)
	}
	temp = temp[0 : len(temp)-1]
}
