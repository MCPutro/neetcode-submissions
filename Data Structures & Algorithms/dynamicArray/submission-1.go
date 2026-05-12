type DynamicArray struct {
    data     []int
	index    int
	capacity int
}

func NewDynamicArray(capacity int) *DynamicArray {
    return &DynamicArray{
		data:     make([]int, capacity),
		index:    0,
		capacity: capacity,
	}
}

func (da *DynamicArray) Get(i int) int {
    return da.data[i]
}

func (da *DynamicArray) Set(i int, n int) {
    da.data[i] = n
}

func (da *DynamicArray) Pushback(n int) {
    if da.index == da.capacity {
        da.resize()
    }
    da.data[da.index] = n
	da.index++
}

func (da *DynamicArray) Popback() int {
    da.index--
    return da.data[da.index]
}

func (da *DynamicArray) resize() {
    temp := make([]int, da.capacity*2)
	copy(temp, da.data)
	da.data = temp
	da.capacity = da.capacity * 2
}

func (da *DynamicArray) GetSize() int {
    return da.index
}

func (da *DynamicArray) GetCapacity() int {
    return da.capacity
}
