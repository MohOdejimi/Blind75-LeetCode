/*  

	128. Longest Consecutive Sequence
Solved
Medium
Topics
premium lock icon
Companies
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.

Example 1:

Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
Example 2:

Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9
Example 3:

Input: nums = [1,0,1,2]
Output: 3
 

Constraints:

0 <= nums.length <= 105
-109 <= nums[i] <= 109
*/

package main 

func maxFN(a, b int) int {
    if a > b {
        return a 
    }
    return b
}

func longestConsecutive(nums []int) int {
    length := 0 
    numsSet := make(map[int]bool)

    for _, num := range nums {
        numsSet[num] = true 
    }

    for _, num := range nums {
        counter := 1 
        nextNum := num + 1
        prevNum := num - 1

        if numsSet[prevNum] {
            continue
        } else {
            for numsSet[nextNum] {
                delete(numsSet, nextNum)
                nextNum++
                counter++
            }
        }
        length = maxFN(counter, length)
    }

    return length
}