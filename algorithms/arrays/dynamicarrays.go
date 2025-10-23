package dynamicarrays

type DynamicArray struct {
	Capacity int
	Size     int
	Data     []int
}

func NewDynamicArray(Capacity int) *DynamicArray {
	return &DynamicArray{
		Capacity: Capacity,
		Size:     0,
		Data:     make([]int, 0, Capacity),
	}
}

func (array *DynamicArray) Get(index int) int {
	return array.Data[index]
}

func (array *DynamicArray) GetSize() int {
	return array.Size
}

func (array *DynamicArray) Pushback(element int) {
	if array.Size == array.Capacity {
		array.Resize()
	}

	array.Data = append(array.Data, element)
	array.Size++
}

func (array *DynamicArray) Popback() int {
	lastElement := array.Data[array.Size-1]

	array.Data = array.Data[:array.Size-1]
	array.Size--

	return lastElement
}

func (array *DynamicArray) Resize() {
	array.Capacity *= 2

	newData := make([]int, array.Size, array.Capacity)
	copy(newData, array.Data)
	array.Data = newData
}

func (array *DynamicArray) GetCapacity() int {
	return array.Capacity
}

func (array *DynamicArray) Set(i int, n int) {
	array.Data[i] = n
}