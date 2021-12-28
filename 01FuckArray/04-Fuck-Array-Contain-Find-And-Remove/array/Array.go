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

// 04 contain find and remove

// 是否包含item元素
func (a *Array) Contains(item int) bool {
	flag := true
	for i := 0; i < a.size; i++ {
		if a.Get(i) == item {
			flag = true
		}
	}
	return flag
}

// 查找数组中元素item所在的索引，不存在则返回 -1
func (a *Array) Find(item int) int {
	index := -1
	for i := 0; i <= a.size; i++ {
		if a.Get(i) == item {
			index = i
		}
	}
	return index
}

// 删除index位置的元素，返回值
func (a *Array) Remove(index int) int {
	if index < 0 || index >= a.size {
		panic("remove failed, index out of range")
	}

	e := a.data[index]
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}
	// 如果如上移动后，最后一位元素还是存在，这里直接赋值
	a.size--
	a.data[a.size] = 0

	return e
}

func (a *Array) RemoveFirst() int {
	return a.Remove(0)
}

func (a *Array) RemoveLast() int {
	return a.Remove(a.size - 1)
}

func (a *Array) RemoveItem(item int) bool {
	flag := false
	index := a.Find(item)
	if index != -1 {
		a.Remove(index)
		flag = true
	}
	return flag
}

// 删除全部 我这里直接new一个新数组
func (a *Array) RemoveAll() {
	a.data = make([]int, a.GetCapacity())
	a.size = 0
}
