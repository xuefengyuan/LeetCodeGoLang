package main

import "fmt"

// 155 最小栈
// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
// 实现 MinStack 类:
//
// MinStack() 初始化堆栈对象。
// void push(int val) 将元素val推入堆栈。
// void pop() 删除堆栈顶部的元素。
// int top() 获取堆栈顶部的元素。
// int getMin() 获取堆栈中的最小元素。
func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println("Min:", minStack.GetMin()) // 返回 -3
	minStack.Pop()
	fmt.Println("Top:", minStack.Top())    // 返回 0
	fmt.Println("Min:", minStack.GetMin()) // 返回 -2
}

type MinStack struct {
	stack    []int // 主栈
	minStack []int // 辅助栈，用于存储最小值
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	// 将元素压入主栈
	this.stack = append(this.stack, val)
	// 如果辅助栈为空，或者当前值小于等于辅助栈的栈顶元素，则压入辅助栈
	if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	// 弹出主栈的栈顶元素
	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	// 如果弹出的元素等于辅助栈的栈顶元素，则同时弹出辅助栈的栈顶元素
	if top == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		panic("stack is empty")
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.minStack) == 0 {
		panic("min stack is empty")
	}
	return this.minStack[len(this.minStack)-1]
}
