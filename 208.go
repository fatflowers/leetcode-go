package main

import (
	"fmt"
)

type Node struct {
	Value uint8
	Children map[uint8]*Node
}

type Trie struct {
	Root Node
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{Node{0, make(map[uint8]*Node)}}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	node := &this.Root
	for i := range word {
		if _, ok := node.Children[word[i]]; !ok {
			if i == len(word) - 1 {
				node.Children[word[i]] = nil
			} else {
				node.Children[word[i]] = &Node{word[i], make(map[uint8]*Node)}
			}
		}

		node = node.Children[word[i]]
	}
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := &this.Root
	for i := range word {
		if node == nil {
			return false
		}

		if _, ok := node.Children[word[i]]; !ok {
			return false
		}
		node = node.Children[word[i]]
	}

	return node == nil
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := &this.Root
	for i := range prefix {
		if node == nil {
			return false
		}

		if _, ok := node.Children[prefix[i]]; !ok {
			return false
		}
		node = node.Children[prefix[i]]
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

//["Trie","insert","search","search","startsWith","startsWith","insert","search","startsWith","insert","search","startsWith"]
//[[],["ab"],["abc"],["ab"],["abc"],["ab"],["ab"],["abc"],["abc"],["abc"],["abc"],["abc"]]

func main() {
	obj := Constructor();
	for _, s := range []string{"Trie","insert","search","search","startsWith","startsWith","insert","search","startsWith","insert","search","startsWith"} {
		obj.Insert(s)
	}

	fmt.Println(obj.Search("abc"))
	fmt.Println(obj.StartsWith("abc"))
	fmt.Println(obj.Search("ab"))
	fmt.Println(obj.StartsWith("ab"))
	fmt.Println(obj.Search(""))
	fmt.Println(obj.StartsWith(""))
}
