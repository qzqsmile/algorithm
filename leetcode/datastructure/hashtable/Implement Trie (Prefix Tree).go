type TrieNode struct{
	node map[rune] *TrieNode
	isWord bool
}

type Trie struct {
	root *TrieNode
}


func Constructor() Trie {
    root := Trie{root:&TrieNode{node: map[rune]*TrieNode{}, isWord:false}}
	return root
}


func (this *Trie) Insert(word string)  {
    node := this.root
    for _, value := range word{
        if node.node[value] == nil{
            node.node[value] = &TrieNode{node: map[rune]*TrieNode{}, isWord: false}
        }
        node = node.node[value]
    }
    node.isWord = true
}


func (this *Trie) Search(word string) bool {
    node := this.root
    for _, value := range word{
        if _, ok := node.node[value]; !ok{
            return false
        }
        node = node.node[value]
    }
    if node.isWord{
        return true
    }
    return false
}


func (this *Trie) StartsWith(prefix string) bool {
    node := this.root 
    for _, value := range prefix{
        if _, ok := node.node[value]; !ok{
            return false
        }
        node = node.node[value]
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */