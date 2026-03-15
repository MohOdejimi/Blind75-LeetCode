

/* Given a string s, find the length of the longest substring without duplicate characters.

 

Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
 

Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces. */

package main

func maxFn(a, b int) int {
    if a > b {
        return a 
    }
    return b
}

func lengthOfLongestSubstring(s string) int {
    sMap := make(map[byte]int)
    result := 0
    n := len(s)

    left := 0 
    
    for right := 0; right < n; right++ {
        ch := s[right]
        sMap[ch]++ 

        if sMap[ch] > 1 {
            for sMap[ch] > 1 {
                sMap[s[left]]--
                left++
            }
        }
        windowSize := right - left + 1 
        result = maxFn(result, windowSize)
    }

    return result 
}