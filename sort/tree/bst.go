package main

import (
	"bytes"
	"data_strcut/comparator"
	"fmt"
	"math/rand"
	"time"
)

type BiSearchTree struct {
	Value  comparator.Comparator
	LChild *BiSearchTree
	RChild *BiSearchTree
}

func (bst *BiSearchTree) writeToBuffer(bf *bytes.Buffer) {
	if bst == nil {
		bf.WriteString("nil")
		return
	}
	if bst.LChild != nil {
		bst.LChild.writeToBuffer(bf)
		bf.WriteByte(' ')
	}
	bf.WriteString(fmt.Sprint(bst.Value))
	if bst.RChild != nil {
		bf.WriteByte(' ')
		bst.RChild.writeToBuffer(bf)
	}
}

func (bst *BiSearchTree) String() string {
	var bf bytes.Buffer
	bf.WriteByte('[')
	bst.writeToBuffer(&bf)
	bf.WriteByte(']')
	return bf.String()
}

func newBiSearchTree() *BiSearchTree {
	return nil
}

func (bst *BiSearchTree) InsertAll(values ...interface{}) *BiSearchTree {
	b := bst
	for i := range values {
		b = b.Insert(values[i])
	}
	return b
}

func (bst *BiSearchTree) Insert(i interface{}) *BiSearchTree {
	c := comparator.Convert(i)
	node := &BiSearchTree{Value: c}
	if bst == nil {
		return node
	}
	b := bst
	for {
		if b.Value.CompareTo(node.Value) <= 0 {
			if b.LChild != nil {
				b = b.LChild
			} else {
				b.LChild = node
				break
			}
		} else {
			if b.RChild != nil {
				b = b.RChild
			} else {
				b.RChild = node
				break
			}
		}
	}
	return bst
}

func (bst *BiSearchTree) Delete(i interface{}) *BiSearchTree {
	return bst.delete(comparator.Convert(i))
}

func (bst *BiSearchTree) delete(c comparator.Comparator) *BiSearchTree {
	if bst == nil {
		return nil
	}
	if i := bst.Value.CompareTo(c); i == 0 {
		if bst.LChild == nil && bst.RChild == nil {
			return nil
		}
		if bst.LChild == nil && bst.RChild != nil {
			return bst.RChild
		}
		if bst.LChild != nil && bst.RChild == nil {
			return bst.LChild
		}
		bst.LChild.MaxNode().RChild = bst.RChild
		bst.RChild = nil
		return bst.LChild.delete(c)
	} else if i < 0 {
		bst.LChild = bst.LChild.delete(c)
	} else {
		bst.RChild = bst.RChild.delete(c)
	}
	return bst
}

func (bst *BiSearchTree) MaxNode() *BiSearchTree {
	if bst == nil {
		return nil
	}
	if bst.RChild != nil {
		return bst.RChild.MaxNode()
	} else {
		return bst
	}
}

func (bst *BiSearchTree) Query(c comparator.Comparator) bool {
	if bst == nil {
		return false
	}
	if i := bst.Value.CompareTo(c); i == 0 {
		return true
	} else if i < 0 {
		return bst.LChild.Query(c)
	} else {
		return bst.RChild.Query(c)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var arr []interface{}
	for i := 0; i < 100000; i++ {
		arr = append(arr, rand.Intn(1000))
	}
	bst := newBiSearchTree()
	fmt.Println(arr)
	bst = bst.InsertAll(arr...)
	//bst = bst.Delete(999)
	fmt.Println(bst)
}
