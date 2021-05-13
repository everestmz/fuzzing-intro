package add

import (
  "bytes"
  "encoding/binary"
)

func FuzzAdd(x, y int) {
  // First, check for commutativity
  if Add(x, y) != Add(y, x) {
    panic("Add not commutative!")
  }

  // Check for associativity
  if Add(Add(x, y), y) != Add(x, Add(y, y)) {
    panic("Add not associative!")
  }
}

// Boilerplate for go-fuzz compatibility
func Fuzz(Data []byte) int {
  buf := bytes.NewBuffer(Data)
  x, err := binary.ReadVarint(buf)
  if err != nil {
    return -1
  }
  y, err := binary.ReadVarint(buf)
  if err != nil {
    return -1
  }
  FuzzAdd(int(x), int(y))
  return 0
}