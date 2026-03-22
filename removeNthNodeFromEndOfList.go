package main 

/* 
Given the head of a linked list, remove the nth node from the end of the list and return its head.

Example 1:


Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
Example 2:

Input: head = [1], n = 1
Output: []
Example 3:

Input: head = [1,2], n = 1
Output: [1]
 

Constraints:

The number of nodes in the list is sz.
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz */


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * } 
 */
func reverseFN(head *ListNode) *ListNode {
    curr := head 
    var  prev *ListNode 

    for curr != nil {
        next := curr.Next 
        curr.Next = prev
        prev = curr
        curr = next
    }

    return prev 
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    reversedHead := reverseFN(head)
    if n == 1 {
        return reverseFN(reversedHead.Next)
    }

    count := 1
    curr := reversedHead 
    var prev *ListNode

    for curr != nil {
        if count == n {
            prev.Next = curr.Next 
            return reverseFN(reversedHead)
        }
        prev = curr 
        curr = curr.Next 
        count++
    } 

    return reverseFN(curr)
}