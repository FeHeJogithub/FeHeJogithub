package main

import "fmt"

// Definición del nodo de la lista enlazada
type ListNode struct {
    Val  int
    Next *ListNode
}

// Función que mezcla dos listas ordenadas
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    tail := dummy

    p1, p2 := list1, list2

    for p1 != nil && p2 != nil {
        if p1.Val < p2.Val {
            tail.Next = p1
            p1 = p1.Next
        } else {
            tail.Next = p2
            p2 = p2.Next
        }
        tail = tail.Next
    }

    if p1 != nil {
        tail.Next = p1
    } else {
        tail.Next = p2
    }

    return dummy.Next
}

// Función para crear una lista desde un slice
func buildList(nums []int) *ListNode {
    if len(nums) == 0 {
        return nil
    }

    head := &ListNode{Val: nums[0]}
    current := head

    for _, v := range nums[1:] {
        current.Next = &ListNode{Val: v}
        current = current.Next
    }

    return head
}

// Función para imprimir una lista enlazada
func printList(head *ListNode) {
    for head != nil {
        fmt.Print(head.Val, " -> ")
        head = head.Next
    }
    fmt.Println("nil")
}

func main() {
    list1 := buildList([]int{1, 3, 4})
    list2 := buildList([]int{1, 2, 5})

    fmt.Println("Lista 1:")
    printList(list1)

    fmt.Println("Lista 2:")
    printList(list2)

    merged := mergeTwoLists(list1, list2)
    fmt.Println("Lista combinada:")
    printList(merged)
}
