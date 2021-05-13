package heap

import "fmt"

func FuzzHeap(ints []int) {
  // First initialize the heap
  heap := NewMinHeap(len(ints))

  // Set up our minimum value
  min := int(^uint(0) >> 1) // The largest int possible in Go

  // Then, insert each int into the heap
  for _, i := range ints {
    err := heap.Insert(i)
    if err != nil {
      // Heap full
      break
    }
    if i < min {
      min = i
    }
  }
  // Also get a second minimum
  // secondMin := int(^uint(0) >> 1)
  // for _, i := range ints {
  //   if i < secondMin && i != min {
  //     secondMin = i
  //   }
  // }

  // Finally, pop an int off the heap to see if the heap is correct
  result := heap.Remove()
  if result != min {
    panic(fmt.Sprintf("Heap result %d != min %d for ints %s", result, min, ints))
  }
  // result = heap.Remove()
  // if result != secondMin {
  //   panic(fmt.Sprintf("Heap result %d != secondMin %d for list %s", result, secondMin, ints))
  // }
}

// Boilerplate to convert from bytes that go-fuzz provides to us
func Fuzz(Data []byte) int {
  if len(Data) < 3 {
    return -1
  }
  ints := []int{}
  for _, b := range Data {
    ints = append(ints, int(b))
  }
  FuzzHeap(ints)
  return 0
}