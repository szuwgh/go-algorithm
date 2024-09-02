package skiplist

import (
	"fmt"
	"testing"
)

func Test_skiplist(t *testing.T) {
	fmt.Println("test1")
	sl := newSkipList()
	sl.insert([]byte("a"), []byte("1"))
	sl.insert([]byte("d"), []byte("4"))
	sl.insert([]byte("c"), []byte("3"))
	sl.insert([]byte("e"), []byte("5"))
	sl.insert([]byte("b"), []byte("2"))
	sl.insert([]byte("f"), []byte("6"))
	sl.insert([]byte("h"), []byte("7"))
	sl.insert([]byte("g"), []byte("10"))
	sl.insert([]byte("j"), []byte("9"))
	sl.insert([]byte("i"), []byte("8"))
	sl.insert([]byte("l"), []byte("11"))
	sl.insert([]byte("k"), []byte("13"))
	sl.insert([]byte("m"), []byte("16"))
	fmt.Println("maxlevel", sl.level)
	sl.Print()
	fmt.Println(string(sl.search([]byte("g")).value))
	// for e := sl.Front(); e != nil; e = e.Next() {
	// 	fmt.Print(string(e.key), "-", string(e.value), " ")
	// }
}
