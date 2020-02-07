package lru

import "fmt"

type Node struct {
	key string
	val interface{}
	pre *Node
	next *Node
}

type LRU struct {
	dataMap map[string]*Node
	head *Node
	tail *Node
	capacity int
	count int
}

func NewLRU(capacity int) *LRU{
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.pre = head
	return &LRU{head:head,
				tail:tail,
				capacity:capacity,
				count:0,
				dataMap:make(map[string]*Node)}
}

func (l *LRU)Get(key string) interface{}{
	v, ok := l.dataMap[key]
	if !ok{
		// 可以从数据库中拉取，然后删除尾节点
		return nil
	}
	// 删除节点，然后至于链表头
	l.detachNode(v)
	l.insertFront(v)
	return v.val
}

// 删除节点
func (l *LRU)detachNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

// 插入头节点
func (l *LRU)insertFront(node *Node){
	node.next = l.head.next
	l.head.next.pre = node
	l.head.next = node
	node.pre = l.head
}

// 删除尾节点
func (l *LRU)delLast(){
	tmp := l.tail.pre // l.tail 指向head，故而删除的是l.tail.pre
	fmt.Println("key",tmp.key)
	tmp.pre.next = l.tail
	l.tail.pre = tmp.pre
	tmp.next = nil
	tmp.pre = nil
	l.count = l.count-1
	delete(l.dataMap,tmp.key)
}

// 加入LRU
func (l *LRU)Set(key string,val interface{}){
	v,ok :=l.dataMap[key]
	if !ok {
		node :=&Node{key:key, val:val}
		if l.count==l.capacity{
			l.delLast()
		}
		l.dataMap[key] = node
		l.insertFront(node)
		l.count = l.count+1
	}else {
		l.detachNode(v)
		l.insertFront(v)
	}
}

