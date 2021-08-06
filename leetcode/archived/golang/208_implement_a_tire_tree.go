//package main
//
//import "fmt"
//
//type TrieNode struct{
//	node map[string] *TrieNode
//	isWord bool
//}
//
//type Trie struct {
//	root *TrieNode
//}
//
//
///** Initialize your data structure here. */
//func Constructor() Trie {
//	root := Trie{root:&TrieNode{node: map[string]*TrieNode{}, isWord:false}}
//	return root
//}
//
//
///** Inserts a word into the trie. */
//func (this *Trie) Insert(word string)  {
//	node := this.root
//	for _, value := range word{
//		if node.node[string(value)] == nil {
//			node.node[string(value)] = &TrieNode{node: map[string]*TrieNode{}, isWord: false}
//		}
//		node = node.node[string(value)]
//	}
//	node.isWord = true
//}
//
//
///** Returns if the word is in the trie. */
//func (this *Trie) Search(word string) bool {
//	node := this.root
//	for _, value := range word{
//		if node.node[string(value)] == nil{
//			return false
//		}
//		node = node.node[string(value)]
//	}
//	return node.isWord
//}
//
//
///** Returns if there is any word in the trie that starts with the given prefix. */
//func (this *Trie) StartsWith(prefix string) bool {
//	node := this.root
//	for _, value := range prefix{
//		if node.node[string(value)] == nil{
//			return false
//		}
//		node = node.node[string(value)]
//	}
//	return true
//}
//
//
///**
// * Your Trie object will be instantiated and called as such:
// * obj := Constructor();
// * obj.Insert(word);
// * param_2 := obj.Search(word);
// * param_3 := obj.StartsWith(prefix);
// */
//
//func main()  {
//	obj := Constructor1()
//	word := "abc"
//	prefix := "acb"
//	obj.Insert(word)
//	param_2 := obj.Search(word)
//	param_3 := obj.StartsWith(prefix)
//	fmt.Println(param_2, param_3)
//
//}
//
