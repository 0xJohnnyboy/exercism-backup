package linkedlist
import "errors"

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type Node struct {
    Value interface{}
    next *Node
    prev *Node
}
type List struct {
    head *Node
    tail *Node
}

func NewList(elements ...interface{}) *List {
    var l List

	for _, v := range elements {
        l.Push(v)
    }

    return &l
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	node := &Node{Value: v}
    if l.head == nil {
        l.head = node
        l.tail = node
        return
    }
    node.next = l.head
    l.head.prev = node
    l.head = node
    return
}

func (l *List) Push(v interface{}) {
	node := &Node{Value: v}
    if l.head == nil {
        l.head = node
        l.tail = node
        return
    }
    node.prev = l.tail
    l.tail.next = node
    l.tail = node
    return
}

func (l *List) Shift() (interface{}, error) {
    if l.head == nil {
        return nil, errors.New("can't shift")
    }
    out := l.head.Value
	l.head = l.head.next
    if l.head == nil {
        l.tail = nil
    } else {
    	l.head.prev = nil
    }

    return out, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.tail == nil {
        return nil, errors.New("can't pop")
    }
    
    out := l.tail.Value
    l.tail = l.tail.prev
    if l.tail == nil {
        l.head = nil 
    } else {
    	l.tail.next = nil
    }
    return out, nil
}

func (l *List) Reverse() {
	node := l.head
    var nextInList *Node
    l.head, l.tail = l.tail, l.head
    for node !=nil {
        nextInList = node.next
        node.next, node.prev = node.prev, node.next
        node = nextInList
    }
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
    return l.tail
}

func (l *List) GetValues() []interface{} {
    var values []interface{}
    node := l.head
    for node != nil {
        values = append(values, node.Value)
        node = node.next
    }
    return values
}
