package heap

import "fmt"

type MinHeap struct {
    heapArray []int
    size      int
    maxsize   int
}

func NewMinHeap(maxsize int) *MinHeap {
    minheap := &MinHeap{
        heapArray: []int{},
        size:      0,
        maxsize:   maxsize,
    }
    return minheap
}

// Public Functions
func (m *MinHeap) Remove() int {
    top := m.heapArray[0]
    m.heapArray[0] = m.heapArray[m.size-2] // Bug here! m.size-1
    m.heapArray = m.heapArray[:(m.size)-1]
    m.size--
    m.downHeapify(0)
    return top
}

func (m *MinHeap) Insert(item int) error {
    if m.size >= m.maxsize {
        return fmt.Errorf("Heal is ful")
    }
    m.heapArray = append(m.heapArray, item)
    m.size++
    m.upHeapify(m.size - 1)
    return nil
}

// Private Functions
func (m *MinHeap) leaf(index int) bool {
    if index >= (m.size/2) && index <= m.size {
        return true
    }
    return false
}

func (m *MinHeap) parent(index int) int {
    return (index - 1) / 2
}

func (m *MinHeap) leftchild(index int) int {
    return 2*index + 1
}

func (m *MinHeap) rightchild(index int) int {
    return 2*index + 2
}

func (m *MinHeap) swap(first, second int) {
    temp := m.heapArray[first]
    m.heapArray[first] = m.heapArray[second]
    m.heapArray[second] = temp
}

func (m *MinHeap) upHeapify(index int) {
    for m.heapArray[index] < m.heapArray[m.parent(index)] {
        m.swap(index, m.parent(index))
        index = m.parent(index)
    }
}

func (m *MinHeap) downHeapify(current int) {
    if m.leaf(current) {
        return
    }
    smallest := current
    leftChildIndex := m.leftchild(current)
    rightRightIndex := m.rightchild(current)
    //If current is smallest then return
    if leftChildIndex < m.size && m.heapArray[leftChildIndex] < m.heapArray[smallest] {
        smallest = leftChildIndex
    }
    if rightRightIndex < m.size && m.heapArray[rightRightIndex] < m.heapArray[smallest] {
        smallest = rightRightIndex
    }
    if smallest != current {
        m.swap(current, smallest)
        m.downHeapify(smallest)
    }
    return
}
func (m *MinHeap) buildMinHeap() {
    for index := ((m.size / 2) - 1); index >= 0; index-- {
        m.downHeapify(index)
    }
}