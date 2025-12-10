package structures

type TrieNode struct {
	IsEnd    bool
	Value    string
	Children map[string]*TrieNode
	Num      int64
}

func NewTrieNode(value string) *TrieNode {
	tn := &TrieNode{
		IsEnd:    false,
		Value:    value,
		Children: nil,
		Num:      -1,
	}
	return tn
}

func (tn *TrieNode) Has(key string) bool {
	if tn.Children == nil {
		return false
	}
	if _, ok := tn.Children[key]; !ok {
		return false
	}
	return true
}

func (tn *TrieNode) Set(key string, val *TrieNode, overwrite bool) {
	if tn.Children == nil {
		tn.Children = make(map[string]*TrieNode)
	}
	if tn.Has(key) && !overwrite {
		return
	}
	tn.Children[key] = val
}

func (tn *TrieNode) Get(key string) *TrieNode {
	if !tn.Has(key) {
		return nil
	}
	return tn.Children[key]
}

func (tn *TrieNode) Delete(key string) {
	if tn.Has(key) {
		delete(tn.Children, key)
	}
	if len(tn.Children) == 0 {
		tn.Children = nil
	}
}
