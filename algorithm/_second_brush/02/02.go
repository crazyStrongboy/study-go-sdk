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

func generateParenthesis(n int) []string {
	t := &T{}
	t.backtrack(n*2, []byte{})
	return t.result
}

var kuohao = []byte{'(', ')'}

type T struct {
	result []string
}

func (t *T) backtrack(n int, track []byte) {
	if n == 0 {
		if isParenthesis(string(track)) {
			t.result = append(t.result, string(track))
		}
		return
	}
	for j := 0; j <= 1; j++ {
		track = append(track, kuohao[j])
		t.backtrack(n-1, track)
		track = track[:len(track)-1]
	}

}

func isParenthesis(s string) bool {
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			cnt++
		} else {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}
