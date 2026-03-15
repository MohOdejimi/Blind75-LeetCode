package main 

/* Given a string s, return the longest palindromic substring in s.

 

Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"
 

Constraints:

1 <= s.length <= 1000
s consist of only digits and English letters. */ 

func longestPalindrome(s string) string {
    n := len(s)
    if n < 2 {
        return s
    }

    length := 1
    str := s[0:1]

    for i := 0; i < n; i++ {
        left, right := i, i
        for left >= 0 && right < n && s[left] == s[right] {
            if right-left+1 > length {
                length = right - left + 1
                str = s[left : right+1]
            }
            left--
            right++
        }

        left, right = i, i+1
        for left >= 0 && right < n && s[left] == s[right] {
            if right-left+1 > length {
                length = right - left + 1
                str = s[left : right+1]
            }
            left--
            right++
        }
    }

    return str
}
