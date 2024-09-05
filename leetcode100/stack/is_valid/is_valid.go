package isValid

// 有效的括号
// 简单
// 相关标签
// 相关企业
// 提示
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

// 有效字符串需满足：

// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。

// 示例 1：

// 输入：s = "()"

// 输出：true

// 示例 2：

// 输入：s = "()[]{}"

// 输出：true

// 示例 3：

// 输入：s = "(]"

// 输出：false

// 示例 4：

// 输入：s = "([])"

// 输出：true

// 代码解释：
// stack：用于存储左括号。
// brackets：映射每个右括号对应的左括号。
// 遍历字符串 s：
// 如果是右括号，检查栈顶元素是否与其匹配，如果匹配则弹出栈顶，否则返回 false。
// 如果是左括号，将其压入栈。
// 最后判断栈是否为空，若为空则表示所有括号都匹配。
// 示例：
// 输入 "()"，输出 true。
// 输入 "()[]{}"，输出 true。
// 输入 "(]"，输出 false。
// 输入 "([])"，输出 true。
func isValid(s string) bool {
	stack := []rune{}
	// 括号的匹配规则
	brackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		// 如果是右括号
		if leftBracket, ok := brackets[char]; ok {
			// 检查栈是否为空或栈顶是否匹配
			if len(stack) == 0 || stack[len(stack)-1] != leftBracket {
				return false
			}
			// 匹配成功则弹出栈顶
			stack = stack[:len(stack)-1]
		} else {
			// 如果是左括号则入栈
			stack = append(stack, char)
		}
	}

	// 栈为空则表示括号匹配有效
	return len(stack) == 0
}
