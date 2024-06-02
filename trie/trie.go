package trie

type Trie struct {
	Children []*Trie
	IsWord   bool
}

func NewTrie() Trie {
	Children := make([]*Trie, 26)
	return Trie{Children, false}
}

func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		this.IsWord = true
		return
	}
	if this.Children[word[0]-'a'] == nil {
		newTrie := NewTrie()
		this.Children[word[0]-'a'] = &newTrie
		newTrie.Insert(word[1:])
	} else {
		this.Children[word[0]-'a'].Insert(word[1:])
	}
}

func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return this.IsWord
	}
	if this.Children[word[0]-'a'] == nil {
		return false
	} else {
		return this.Children[word[0]-'a'].Search(word[1:])
	}
}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	if this.Children[prefix[0]-'a'] == nil {
		return false
	} else {
		return this.Children[prefix[0]-'a'].StartsWith(prefix[1:])
	}
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
