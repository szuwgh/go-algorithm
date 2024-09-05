package decode_str

// 字符串解码

// 给定一个经过编码的字符串，返回它解码后的字符串。

// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

// 示例 1：

// 输入：s = "3[a]2[bc]"
// 输出："aaabcbc"
// 示例 2：

// 输入：s = "3[a2[c]]"
// 输出："accaccacc"
// 示例 3：

// 输入：s = "2[abc]3[cd]ef"
// 输出："abcabccdcdcdef"
// 示例 4：

// 输入：s = "abc3[cd]xyz"
// 输出："abccdcdcdxyz"

// 解码字符串的问题可以通过使用栈（stack）来解决。
// 我们可以将遇到的数字、字符串和方括号依次处理，使用栈的特性来解码嵌套结构。具体思路如下：

// 解题思路：
// 遇到数字：将数字解析出来（可能是多位数），表示接下来要重复的次数。
// 遇到 [：说明一个新的重复序列开始，将当前的字符串和数字压入栈中，开始记录新的字符序列。
// 遇到 ]：说明当前的重复序列结束，弹出栈顶的数字和之前的字符串，将当前的字符串重复相应的次数后，拼接回去。
// 遇到字母：直接添加到当前的临时字符串中。

func decodeString(s string) string {
	// 栈，用来存储之前的结果和数字
	stack := []string{}
	numStack := []int{}
	currentString := ""
	num := 0

	for i := 0; i < len(s); i++ {
		ch := s[i]

		// 如果是数字
		if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0') // 处理多位数字
		} else if ch == '[' {
			// 遇到左括号时，将当前数字和字符串推入栈中，开始新的字符串记录
			numStack = append(numStack, num)
			stack = append(stack, currentString)
			// 重置当前记录
			currentString = ""
			num = 0
		} else if ch == ']' {
			// 遇到右括号时，弹出栈顶数字，并将当前字符串重复相应次数
			prevString := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			repeatTimes := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			// 将当前字符串重复指定次数
			temp := ""
			for j := 0; j < repeatTimes; j++ {
				temp += currentString
			}
			// 拼接之前的字符串
			currentString = prevString + temp
		} else {
			// 直接添加字母
			currentString += string(ch)
		}
	}

	return currentString
}
