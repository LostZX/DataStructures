package array

type Array struct {
	data []int
	size int
}

func Instance(capacity int) *Array {
	return &Array{data: make([]int, capacity)}
}

func (arr *Array) GetCapacity() int {
	return len(arr.data)
}

func (arr *Array) GetSize() int {
	return arr.size
}

func (arr *Array) IsEmpty() bool {
	return arr.size == 0
}
