package main 

/* 
	Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
 

Example 1:

Input: s = "()"

Output: true

Example 2:

Input: s = "()[]{}"

Output: true

Example 3:

Input: s = "(]"

Output: false

Example 4:

Input: s = "([])"

Output: true

Example 5:

Input: s = "([)]"

Output: false

 

Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'. */

func isValid(s string) bool {
    stack := []rune{}
    charMap := make(map[rune]rune)
    n := len(s)

    charMap[')'] = '('
    charMap['}'] = '{'
    charMap[']'] = '['

    for i := range n {
        ch := rune(s[i])
        if ch == '(' || ch == '{' || ch == '[' {
            stack = append(stack, ch)
        } else {
            openingBracket := charMap[ch]
            if len(stack) == 0 || openingBracket != stack[len(stack) - 1] {
                return false
            }
            stack = stack[:len(stack) - 1]
        }
    }
    return len(stack) == 0
}