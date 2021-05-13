package multiply

package add

import (
  "bytes"
  "encoding/binary"
)

func FuzzMultiply(x, y, z int) {
  // HINT: the arguments of this function aren't a mistake
  // some of multiply's properties require 3 numbers
  
  // Add your properties here
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
  z, err := binary.ReadVarint(buf)
  if err != nil {
    return -1
  }
  FuzzMultiply(int(x), int(y), int(z))
  return 0
}