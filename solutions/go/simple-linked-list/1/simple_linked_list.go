package linkedlist
import "errors"

// Define the List and Element types here.
type List struct {
    head *Element
    tail *Element
}

type Element struct {
    Id int
    next *Element
}

func New(elements []int) *List {
	list := &List{}
    
    if len(elements) == 0 {
        return list
    }

    for _, v := range elements {
        list.Push(v)
    }

    return list
}

func (l *List) Size() int {
	if l.head == nil {
        return 0
    }
    size := 1
    node := l.head
    
    for node.next != nil {
        size++
        node = node.next
    }
    
    return size
}

func (l *List) Push(element int) {
    newElement := &Element { Id: element }
    
	if l.head == nil {
        l.head = newElement
        l.tail = newElement
        return
    } 
    
	l.tail.next = newElement
    l.tail = newElement
    
    return
}

func (l *List) Pop() (int, error) {
	if l.tail == nil {
        return 0, errors.New("can't pop empty list")
    }

    out := l.tail.Id

    if l.tail == l.head {
        l.tail = nil
        l.head = nil
        return out, nil
    }
    
    node := l.head

    for i:=1; i<l.Size() - 1; i++ {
        node = node.next
    }
    
    node.next = nil
    l.tail = node

    return out, nil
}

func (l *List) Array() []int {
	var array []int
    
    if l.head == nil {
        return array
    }
    
    node := l.head
    
    for node != nil {
        array = append(array, node.Id)
        node = node.next
    }
    
    return array
}

func (l *List) Reverse() *List {
	if l.head == nil {
        return l
    }

    node := l.head
    var prev *Element
    
    for node != nil {
        nextNodeToWalk := node.next
        if prev!=nil {
            node.next = prev
        }

        prev = node
        node = nextNodeToWalk
    }
    
    l.head, l.tail = l.tail, l.head
    l.tail.next = nil
    
    return l
}
