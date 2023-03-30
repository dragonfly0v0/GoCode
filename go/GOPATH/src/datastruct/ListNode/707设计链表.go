/*
在链表类中实现这些功能：

get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
*/
package main

type SingleNode struct {
	val  int
	next *SingleNode
}

type MyLinkedList struct {
	Size        int
	dummpy_node *SingleNode
}

func Constructor() MyLinkedList {
	var node *MyLinkedList
	node.Size = 0
	node.dummpy_node = new(SingleNode)
	return *node
}

func (this *MyLinkedList) Get(index int) int {
	if this == nil || index < 0 || index > this.Size {
		return -1
	}

	cur := this.dummpy_node.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	var head_node *SingleNode
	head_node.val = val

	head_node.next = this.dummpy_node.next
	this.dummpy_node.next = head_node
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	var tail_node *SingleNode
	tail_node.val = val

	cur := this.dummpy_node
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = tail_node
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Size {
		return
	} else if index < 0 {
		index = 0
	}
	var node *SingleNode
	node.val = val

	cur := this.dummpy_node.next

	for i := 0; i < index; i++ {
		cur = cur.next
	}

	node.next = cur.next
	cur.next = node
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index > this.Size || index < 0 {
		return
	}

	cur := this.dummpy_node
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur = cur.next.next
	this.Size--
}
