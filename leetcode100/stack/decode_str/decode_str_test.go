package decode_str

import (
	"fmt"
	"testing"
)

func TestDecodeStr(t *testing.T) {
	// 示例测试
	fmt.Println(decodeString("3[a]2[bc]"))     // 输出: "aaabcbc"
	fmt.Println(decodeString("3[a2[c]]"))      // 输出: "accaccacc"
	fmt.Println(decodeString("2[abc]3[cd]ef")) // 输出: "abcabccdcdcdef"
	fmt.Println(decodeString("abc3[cd]xyz"))   // 输出: "abccdcdcdxyz"
}
