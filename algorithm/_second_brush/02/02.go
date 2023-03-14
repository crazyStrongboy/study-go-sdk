package main

func isValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, ')')
		} else if s[i] == '[' {
			stack = append(stack, ']')
		} else if s[i] == '{' {
			stack = append(stack, '}')
		} else {
			if len(stack) > 0 {
				v := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if v != s[i] {
					return false
				}
			} else {
				return false
			}
		}
		//fmt.Println(m)
	}
	return len(stack) == 0
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	if list1 == nil {
		dummy.Next = list2
	} else if list2 == nil {
		dummy.Next = list1
	} else {
		if list1.Val < list2.Val {
			dummy.Next = &ListNode{Val: list1.Val, Next: mergeTwoLists(list1.Next, list2)}
		} else {
			dummy.Next = &ListNode{Val: list2.Val, Next: mergeTwoLists(list1, list2.Next)}
		}
	}
	return dummy.Next
}
