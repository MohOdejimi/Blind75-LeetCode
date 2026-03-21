package main 

/* 
You are given the head of a singly linked-list. The list can be represented as:

L0 → L1 → … → Ln - 1 → Ln
Reorder the list to be on the following form:

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
You may not modify the values in the list's nodes. Only nodes themselves may be changed.

 

Example 1:


Input: head = [1,2,3,4]
Output: [1,4,2,3]
Example 2:


Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]
 

Constraints:

The number of nodes in the list is in the range [1, 5 * 104].
1 <= Node.val <= 1000 */

type Listnode struct {
      Val int
      Next *Listnode
}
 
func reversedHalf(head *Listnode) *Listnode {
    curr := head  
    var prev *Listnode 

    for curr != nil {
        next := curr.Next 
        curr.Next = prev 
        prev = curr
        curr =  next
    }

    return prev 
}

func median(head *Listnode) *Listnode {
    fast := head 
    slow := head 

    for fast != nil && fast.Next != nil  {
        fast = fast.Next.Next 
        slow = slow.Next
    }

    return slow  
}

func mergeFN(l1, l2 *Listnode) *Listnode {
    for l1 != nil && l2 != nil {
        l1Next := l1.Next 
        l2Next := l2.Next 

        l1.Next = l2 
        l2.Next = l1Next

        l1 = l1Next 
        l2 = l2Next
    }
    return l1
}


func reorderList(head *Listnode)  {
    medianNode := median(head)
    secondHalf := medianNode.Next 
    medianNode.Next =  nil 
    reversedSecondHalf := reversedHalf(secondHalf)

    mergeFN(head, reversedSecondHalf)
}