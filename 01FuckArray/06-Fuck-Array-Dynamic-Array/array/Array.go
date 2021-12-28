package array

import (
	"../utils"
	"bytes"
	"fmt"
)

//太多了就不一个个改了

// 自动扩容

type Array struct {
	data []interface{}
	size int
}

// 构造函数，传入数组的容量capacity构造Array
// 动态数组 就不需要初始化了，给了5的初始值，不够自动扩容
func Instance() *Array {
	return &Array{
		data: make([]interface{}, 5),
	}
}

// 获取数组的容量
func (a *Array) GetCapacity() int {
	return len(a.data)
}

// 获得数组中的元素个数
func (a *Array) GetSize() int {
	return a.size
}

// 返回数组是否为空
func (a *Array) IsEmpty() bool {
	return a.size == 0
}

// 在第 index 个位置插入一个新元素 e
// 新增 自动扩容
func (a *Array) Add(index int, e interface{}) {
	if index < 0 || index > a.size {
		panic("add failed, index out of range")
	}

	if a.size == len(a.data) {
		a.resize()
	}

	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}

	a.data[index] = e
	a.size++
}

// 向所有元素后添加一个新元素
func (a *Array) AddLast(e interface{}) {
	a.Add(a.size, e)
}

// 向所有元素前添加一个新元素
func (a *Array) AddFirst(e interface{}) {
	a.Add(0, e)
}

// 获取 index 索引位置的元素
func (a *Array) Get(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("get failed, index out of range")
	}
	return a.data[index]
}

// 修改 index 索引位置的元素
func (a *Array) Set(index int, e interface{}) {
	if index < 0 || index >= a.size {
		panic("set failed, index out of range")
	}
	a.data[index] = e
}

// 查找数组中是否有元素 e
func (a *Array) Contains(e interface{}) bool {
	for i := 0; i < a.size; i++ {
		if utils.Compare(a.data[i], e) == 0 {
			return true
		}
	}
	return false
}

// 查找数组中元素 e 所在的索引，不存在则返回 -1
func (a *Array) Find(e interface{}) int {
	for i := 0; i < a.size; i++ {
		if utils.Compare(a.data[i], e) == 0 {
			return i
		}
	}
	return -1
}

// 查找数组中元素 e 所有的索引组成的切片，不存在则返回 -1
func (a *Array) FindAll(e interface{}) (indexes []int) {
	for i := 0; i < a.size; i++ {
		if utils.Compare(a.data[i], e) == 0 {
			indexes = append(indexes, i)
		}
	}
	return
}

// 从数组中删除 index 位置的元素，返回删除的元素
func (a *Array) Remove(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("remove failed, index out of range")
	}

	e := a.data[index]
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}
	a.size--
	a.data[a.size] = nil //loitering object != memory leak
	return e
}

// 从数组中删除第一个元素，返回删除的元素
func (a *Array) RemoveFirst() interface{} {
	return a.Remove(0)
}

// 从数组中删除最后一个元素，返回删除的元素
func (a *Array) RemoveLast() interface{} {
	return a.Remove(a.size - 1)
}

// 从数组中删除一个元素 e
func (a *Array) RemoveElement(e interface{}) bool {
	index := a.Find(e)
	if index == -1 {
		return false
	}

	a.Remove(index)
	return true
}

// 从数组中删除所有元素 e
func (a *Array) RemoveAllElement(e interface{}) bool {
	if a.Find(e) == -1 {
		return false
	}

	for i := 0; i < a.size; i++ {
		if utils.Compare(a.data[i], e) == 0 {
			a.Remove(i)
		}
	}
	return true
}

// 重写 Array 的 string 方法
func (a *Array) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", a.size, len(a.data)))
	buffer.WriteString("[")
	for i := 0; i < a.size; i++ {
		// fmt.Sprint 将 interface{} 类型转换为字符串
		buffer.WriteString(fmt.Sprintf("%v", a.data[i]))
		if i != (a.size - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}

// 扩容
func (a *Array) resize() {
	newCapacity := a.GetCapacity() * 2
	newData := make([]interface{}, newCapacity)
	for i := 0; i < a.size; i++ {
		newData[i] = a.data[i]
	}

	a.data = newData
}
