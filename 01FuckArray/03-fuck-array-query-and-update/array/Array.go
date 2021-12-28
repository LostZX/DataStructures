package array

type Array struct {
	// 内部维护的数组
	data []int
	// 当前大小
	size int
}

func Instance(capacity int) *Array {
	// 初始化array size默认为0
	return &Array{data: make([]int, capacity)}
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

func (a *Array) AddFirst(item int) {
	a.Add(0, item)
}

// 在第index个位置插入新元素
func (a *Array) Add(index int, item int) {
	if a.GetCapacity() == a.size {
		panic("add failed, array is full")
	}
	if index < 0 || index > a.size {
		panic("add failed, index out of range")
	}

	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[index] = item
	a.size++
}

func (a *Array) AddLast(item int) {
	a.Add(a.size, item)
}

// 03 query and update

// 获取 index 索引位置的元素
func (a *Array) Get(index int) int {
	// 判断下标是否越界
	if index > a.GetCapacity() || index < 0 {
		panic("get failed, index out of range")
	}

	return a.data[index]
}

func (a *Array) Set(index int, item int) bool {
	if index > a.GetCapacity() || index < 0 {
		return false
	}

	a.data[index] = item
	return true
}
