// tree_impl.go
package tree

import (
	"fmt"
	"github.com/wdsgyj/golf/linear"
)

type Node struct {
	parent     *Node
	brother    *Node
	child      *Node
	depth      int // start from 0，如果 isCopied 为 true，则本变量的值无效
	horizontal int // start from 0，如果 isCopied 为 true，则本变量的值无效
	Value      interface{}
	isCopied   bool
}

func (this *Node) String() string {
	return fmt.Sprintf("[%d, %d]: %v", this.depth, this.horizontal, this.Value)
}

func (this *Node) clear() {
	this.brother = nil
	this.child = nil
	this.parent = nil
	this.Value = nil
}

func (this *Node) Copy() *Node {
	rs := new(Node)
	rs.brother = this.brother
	rs.child = this.child
	rs.depth = this.depth
	rs.horizontal = this.horizontal
	rs.parent = this.parent
	rs.Value = this.Value
	rs.isCopied = true // 标记为 copy 得到
	return rs
}

func (this *Node) MakeSubTree() *Tree {
	rs := new(Tree)
	rs.root = (*Node)(this.Copy())
	rs.root.parent = nil
	return rs
}

func (this *Node) GetRoot() *Node {
	rs := this
	for rs.parent != nil {
		rs = rs.parent
	}
	return rs
}

func (n *Node) AddBrother(v interface{}) *Node {
	parent := n.parent
	lastBrother := parent.child
	for lastBrother.brother != nil {
		lastBrother = lastBrother.brother
	}

	lastBrother.brother = new(Node)
	lastBrother.brother.parent = parent
	lastBrother.brother.depth = lastBrother.depth
	lastBrother.brother.horizontal = lastBrother.horizontal + 1
	lastBrother.brother.Value = v
	return lastBrother.brother
}

func (n *Node) AddChild(v interface{}) *Node {
	if n.child == nil {
		n.child = new(Node)
		n.child.parent = n
		n.child.depth = n.depth + 1
		n.child.Value = v
		return n.child
	} else {
		return n.child.AddBrother(v)
	}
}

func (n *Node) GetParent() *Node {
	return n.parent
}

func (n *Node) GetChild() *Node {
	return n.child
}

func (n *Node) GetBrother() *Node {
	return n.brother
}

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return new(Tree)
}

func (this *Tree) IsEmpty() bool {
	return this.root == nil
}

func (this *Tree) AddRoot(v interface{}) *Node {
	this.root = new(Node)
	this.root.Value = v
	return this.root
}

func (this *Tree) GetRoot() *Node {
	return this.root
}

func (this *Tree) Remove(n *Node) {
	parent := n.parent

	// n 为第一个孩子
	if parent.child == n {
		parent.child = n.brother
	} else {
		lastBrother := parent.child
		for lastBrother.brother != n {
			lastBrother = lastBrother.brother
		}

		// lastBrother 为 n 的左边兄弟节点
		lastBrother.brother = n.brother
	}

	n.clear()
}

func (this *Tree) Interator(fn func(interface{})) {
	if this.root == nil {
		return
	}

	que := linear.NewQueue()
	que.Add(this.root)
	var node *Node

	for !que.IsEmpty() {
		node = que.Remove().(*Node)
		fn(node.Value)
		for n := node.child; n != nil; n = n.brother {
			que.Add(n)
		}
	}
}

// 深度优先遍历递归版本
func (this *Tree) interatorDeepRecursive(fn func(interface{})) {
	this.internalInteratorRecursive(this.root, fn)
}

func (this *Tree) internalInteratorRecursive(n *Node, fn func(interface{})) {
	if n == nil {
		return
	}

	fn(n.Value)

	for node := n.child; node != nil; node = node.brother {
		this.internalInteratorRecursive(node, fn)
	}
}

func (this *Tree) interatorDeep(fn func(interface{})) {
	que := linear.NewStack()
	que.Add(this.root)
	for !que.IsEmpty() {
		node := que.Remove().(*Node)
		if node == nil {
			continue
		}

		fn(node.Value)

	}
}