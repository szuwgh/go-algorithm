package btree

import (
	"bytes"
	"fmt"
)

var MIN_FOR_LEAF int              //叶子节点最小键数
var MAX_FOR_LEAF int              //叶子节点最大键数
var MAX_CHILDREN_FOR_INTERNAL int //内节点的孩子最大数
//var MAX_CHILDREN_FOR_INTERNAL int //内节点的孩子最大数

// B+树实现
type BPlusTree struct {
	root Node
}

func NewBPlusTree(factor int) *BPlusTree {
	MAX_FOR_LEAF = factor - 1
	MAX_CHILDREN_FOR_INTERNAL = factor
	bpTree := &BPlusTree{}
	bpTree.root = NewLeafNode()
	return bpTree
}

func (bt *BPlusTree) Set(k, v []byte) {
	node := bt.root.Insert(k, v)
	if node != nil {
		bt.root = node
	}
}

type Node interface {
	Insert([]byte, []byte) Node
	Get([]byte) []byte
	setParent(Node)
}

//内节点
type InternalNode struct {
	size      int  //大小
	parent    Node //父节点
	keys      [][]byte
	childNode []Node //子节点
}

func NewInternalNode() *InternalNode {
	internalNode := &InternalNode{}
	internalNode.size = 0
	internalNode.keys = make([][]byte, MAX_CHILDREN_FOR_INTERNAL)
	internalNode.childNode = make([]Node, MAX_CHILDREN_FOR_INTERNAL)
	return internalNode
}

func (in *InternalNode) setParent(parent Node) {
	in.parent = parent
}

func (in *InternalNode) print() {
	for _, v := range in.keys {
		fmt.Print(string(v) + " ")
	}
}

//插入
func (in *InternalNode) Insert(k []byte, v []byte) Node {
	i := 0
	for ; i < in.size; i++ {
		if bytes.Compare(k, in.keys[i]) < 0 {
			break
		}
	}
	return in.childNode[i].Insert(k, v)
}

func (in *InternalNode) InsertNode(k []byte, leftChild, rightChild Node) Node {
	if in.size == 0 {
		in.size++
		in.childNode[0] = leftChild
		in.childNode[1] = rightChild
		in.keys[0] = k
		return in
	}
	newKeys := make([][]byte, MAX_CHILDREN_FOR_INTERNAL+1)
	newChildNode := make([]Node, MAX_CHILDREN_FOR_INTERNAL+2)
	i := 0
	for ; i < in.size; i++ {
		n := bytes.Compare(k, in.keys[i])
		if n == -1 {
			break
		}
	}

	copy(newKeys, in.keys[:i])
	newKeys[i] = k
	copy(newKeys, in.keys[i+1:])

	//插入值
	copy(newChildNode, in.childNode[:i])
	newChildNode[i+1] = rightChild
	copy(newChildNode, in.childNode[i+2:])

	in.size++
	if in.size <= MAX_CHILDREN_FOR_INTERNAL {
		in.keys = newKeys
		in.childNode = newChildNode
		return nil
	}
	m := in.size / 2

	//分裂内节点
	//创建一个新的内节点
	newNode := NewInternalNode()
	newNode.size = in.size - m - 1
	copy(newNode.keys, newKeys[m+1:])
	copy(newNode.childNode, newChildNode[m+1:])

	for j := 0; j <= newNode.size; j++ {
		newNode.childNode[j].setParent(newNode)
	}
	in.size = m
	in.keys = make([][]byte, MAX_CHILDREN_FOR_INTERNAL)
	in.childNode = make([]Node, MAX_CHILDREN_FOR_INTERNAL)
	copy(in.keys, newKeys[:m])
	copy(in.childNode, newChildNode[:m])
	//在父节点插入值
	var parent *InternalNode
	if in.parent == nil {
		parent = NewInternalNode()
		in.parent = parent
	}
	newNode.parent = parent
	return parent.InsertNode(newNode.keys[0], in, newNode)
}

func (in *InternalNode) Get(k []byte) []byte {
	return nil
}

//叶子节点
type LeafNode struct {
	size   int
	parent Node
	keys   [][]byte //keys
	values [][]byte //值
}

func NewLeafNode() *LeafNode {
	leafNode := &LeafNode{}
	leafNode.keys = make([][]byte, MAX_FOR_LEAF)
	leafNode.values = make([][]byte, MAX_FOR_LEAF)
	leafNode.parent = nil
	return leafNode
}

func (in *LeafNode) setParent(parent Node) {
	in.parent = parent
}

func (in *LeafNode) print() {
	for _, v := range in.keys {
		fmt.Print(string(v) + " ")
	}

}

func (in *LeafNode) Get(k []byte) []byte {
	return nil
}

//插入
func (ln *LeafNode) Insert(k []byte, v []byte) Node {
	newKeys := make([][]byte, MAX_FOR_LEAF+1)
	newValues := make([][]byte, MAX_FOR_LEAF+1)
	i := 0
	for ; i < ln.size; i++ {
		n := bytes.Compare(k, ln.keys[i])
		if n == 0 {
			ln.values[i] = v
			return nil
		} else if n == -1 {
			break
		}
	}

	//插入键 深拷贝
	//slice1 := []int{1, 2, 3, 4, 5}
	// slice2 := []int{5, 4, 3}
	// copy(slice2, slice1) 只会复制slice1的前3个元素到slice2中
	// copy(slice1, slice2) 只会复制slice2的3个元素到slice1的前3个位置
	copy(newKeys, ln.keys[:i])
	newKeys[i] = k
	copy(newKeys, ln.keys[i+1:])

	//插入值
	copy(newValues, ln.values[:i])
	newValues[i] = v
	copy(newValues, ln.values[i+1:])

	ln.size++
	if ln.size <= MAX_FOR_LEAF {
		ln.keys = newKeys
		ln.values = newValues
		return nil
	}

	//节点分裂
	m := ln.size / 2
	ln.keys = make([][]byte, MAX_FOR_LEAF)
	ln.values = make([][]byte, MAX_FOR_LEAF)
	copy(ln.keys, newKeys[:m])
	copy(ln.values, newValues[:m])
	ln.size = m
	//新的右节点
	newNode := NewLeafNode()
	newNode.size = ln.size - m
	copy(newNode.keys, newKeys[m+1:])
	copy(newNode.keys, newKeys[m+1:])

	//在父节点插入值
	var parent *InternalNode
	if ln.parent == nil {
		parent = NewInternalNode()
		ln.parent = parent
	}
	newNode.parent = parent
	return parent.InsertNode(newNode.keys[0], ln, newNode)
}
