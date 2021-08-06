package main


type TrieNode struct{
	node map[string] *TrieNode
	isWord bool
}

type Trie struct {
	root *TrieNode
}


/** Initialize your data structure here. */
func Constructor1() Trie {
	root := Trie{root:&TrieNode{node: map[string]*TrieNode{}, isWord:false}}
	return root
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	node := this.root
	for _, value := range word{
		if node.node[string(value)] == nil {
			node.node[string(value)] = &TrieNode{node: map[string]*TrieNode{}, isWord: false}
		}
		node = node.node[string(value)]
	}
	node.isWord = true
}


func findWords(board [][]byte, words []string) []string {
	root := Constructor1()
	var res []string
	for _, word := range words{
		root.Insert(word)
	}
	for i := range(board){
		for j := range(board[0]){
			dfs(root.root, i, j, board, "", &res)
		}
	}
	return res
}

func dfs(node *TrieNode, i int, j int, board [][]byte, path string, res *[]string){
	if node.isWord{
		*res = append(*res, path)
		node.isWord = false
	}
	if (i < 0) || (i >= len(board)) || (j < 0) || (j >= len(board[0])){
		return
	}
	tmp := board[i][j]
	if node.node[string(tmp)] == nil {
		return
	}
	board[i][j] = '#'
	dfs(node.node[string(tmp)], i+1, j, board, path+string(tmp), res)
	dfs(node.node[string(tmp)], i-1, j, board, path+string(tmp), res)
	dfs(node.node[string(tmp)], i, j+1, board, path+string(tmp), res)
	dfs(node.node[string(tmp)], i, j-1, board, path+string(tmp), res)
	board[i][j] = tmp
}

//func test(node *[]int){
//	*node = append(*node, 1)
//}
//
//func main()  {
//	node := []int{2,3,4}
//	test(&node)
//	fmt.Println(node)
//}