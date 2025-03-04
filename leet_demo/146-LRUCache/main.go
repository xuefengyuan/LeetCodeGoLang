package main

import "fmt"

// 146 LRU 缓存
// 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
// 实现 LRUCache 类：
// LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
func main() {
	// 创建容量为 2 的 LRU 缓存
	cache := Constructor(2)

	// 测试用例
	cache.Put(1, 1)           // 添加 key=1, value=1
	cache.Put(2, 2)           // 添加 key=2, value=2
	fmt.Println(cache.Get(1)) // 返回 1
	cache.Put(3, 3)           // 添加 key=3, 导致 key=2 被移除
	fmt.Println(cache.Get(2)) // 返回 -1（未找到）
	cache.Put(4, 4)           // 添加 key=4, 导致 key=1 被移除
	fmt.Println(cache.Get(1)) // 返回 -1（未找到）
	fmt.Println(cache.Get(3)) // 返回 3
	fmt.Println(cache.Get(4)) // 返回 4

	// 打印缓存状态
	cache.PrintCache()
}

func (lru *LRUCache) PrintCache() {
	node := lru.list.head.next
	fmt.Print("Cache: ")
	for node != nil && node != lru.list.tail {
		fmt.Printf("[%d:%d] ", node.key, node.value)
		node = node.next
	}
	fmt.Println()
}

// Node 定义双向链表的节点结构
type Node struct {
	// 节点的键和值
	key, value int
	// 双向链表的前驱和后继指针
	prev, next *Node
}

// DoubleLinkedList 定义双向链表
type DoubleLinkedList struct {
	// 头尾虚拟节点
	head, tail *Node
}

// NewDoubleLinkedList 创建一个空的双向链表
func NewDoubleLinkedList() *DoubleLinkedList {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &DoubleLinkedList{head: head, tail: tail}
}

// AddToHead 将节点添加到链表头部
func (dll *DoubleLinkedList) AddToHead(node *Node) {
	node.next = dll.head.next
	node.prev = dll.head
	dll.head.next.prev = node
	dll.head.next = node
}

// RemoveNode 从链表中移除指定节点
func (dll *DoubleLinkedList) RemoveNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// RemoveTail 移除链表尾部节点并返回它
func (dll *DoubleLinkedList) RemoveTail() *Node {
	// 判断头节点的下一个节点等于尾节点，表示链表为空
	if dll.head.next == dll.tail {
		return nil
	}
	tail := dll.tail.prev
	dll.RemoveNode(tail)
	return tail
}

// LRUCache 定义 LRU 缓存结构
type LRUCache struct {
	// 缓存的最大容量
	capacity int
	// 哈希表，存储键和节点的映射
	cache map[int]*Node
	// 双向链表，维护访问顺序
	list *DoubleLinkedList
}

// Constructor 创建 LRU 缓存
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		list:     NewDoubleLinkedList(),
	}
}

// Get 获取缓存中指定 key 的值
func (this *LRUCache) Get(key int) int {
	// 如果 key 不存在，返回 -1
	if node, ok := this.cache[key]; ok {
		// 将节点移动到链表头部
		this.list.RemoveNode(node)
		this.list.AddToHead(node)
		return node.value
	}

	return -1
}

// Put 向缓存中添加或更新一个键值对
func (this *LRUCache) Put(key int, value int) {
	// 如果 key 已存在，更新值并移动到链表头部
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.list.RemoveNode(node)
		this.list.AddToHead(node)
		return
	}

	// 如果 key 不存在，检查是否需要移除最少使用的节点
	if len(this.cache) == this.capacity {
		// 移除链表尾部节点（最近最少使用的节点）
		tail := this.list.RemoveTail()
		delete(this.cache, tail.key)
	}

	// 添加新节点到链表头部
	newNode := &Node{key: key, value: value}
	this.cache[key] = newNode
	this.list.AddToHead(newNode)
}
