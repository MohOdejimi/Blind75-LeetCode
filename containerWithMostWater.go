/*	 
11. Container With Most Water
Medium
Topics
premium lock icon
Companies
Hint
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container. */

// My Solution 

package main 

func minFN(left, right int) int {
    if left < right {
        return left
    }
    return right
}

func maxArea(height []int) int {
    maxArea := 0
    n := len(height)
    

    left := 0
    right := n - 1 

    for left <= right {
        breadth := right - left 
 
        area := breadth * minFN(height[left], height[right])
        if area > maxArea {
            maxArea = area 
        }
        if height[left] < height[right] {
            left++
        } else if height[right] < height[left] {
            right--
        } else {
            left++
            right--
        }
    }


    return maxArea
}