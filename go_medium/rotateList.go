/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func rotateRight(head *ListNode, k int) *ListNode {

	// Create an array of all the pointers

	var out_head *ListNode
	var current_node *ListNode
    var length_of_list int = 0
    var new_head_index int 

    if head == nil || head.Next == nil {return head}

	current_node = head

    // get length of the list and turn it into a circular list
    for current_node.Next != nil {
        current_node = current_node.Next
        length_of_list++
    }
    length_of_list++

    if length_of_list == k {return head}

    current_node.Next = head // create a circular list

    if k > length_of_list {
        new_head_index = length_of_list - k%length_of_list
    } else {
        new_head_index = length_of_list - k
    }
    
    current_node = head;

	// set the new head and the tail the node before
    for i:= 0; i < new_head_index - 1; i++ {
        current_node = current_node.Next
    }

    out_head = current_node.Next
    current_node.Next = nil
    
    return out_head
