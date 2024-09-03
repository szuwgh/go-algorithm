package addtwonum

import "testing"

func TestAddTwoNum(t *testing.T) {
	l1 := createList([]int{2, 4, 3})
	l2 := createList([]int{5, 6, 4})
	result := addTwoNumbers(l1, l2)
	printList(result)

	l1 = createList([]int{0})
	l2 = createList([]int{0})
	result = addTwoNumbers(l1, l2)
	printList(result)

	l1 = createList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 = createList([]int{9, 9, 9, 9})
	result = addTwoNumbers(l1, l2)
	printList(result)
}
