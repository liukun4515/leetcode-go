package tests

import (
	"testing"
)

/**
 * [141] Linked List Cycle
 *
 * Given a linked list, determine if it has a cycle in it.
 *
 * To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: head = 1
 * Output: true
 * Explanation: There is a cycle in the linked list, where tail connects to the second node.
 *
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: head = 0
 * Output: true
 * Explanation: There is a cycle in the linked list, where tail connects to the first node.
 *
 *
 *
 *
 *
 * Example 3:
 *
 *
 * Input: head = -1
 * Output: false
 * Explanation: There is no cycle in the linked list.
 *
 *
 *
 *
 *
 *
 * Follow up:
 *
 * Can you solve it using O(1) (i.e. constant) memory?
 *
 */

func TestLinkedListCycle(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2, Next: node1}
	node1.Next = node2
	if hasCycle(node1) != true {
		t.Fail()
	}
}

// submission codes start here

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for {
		if fast == nil {
			return false
		}
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
}

// submission codes end
