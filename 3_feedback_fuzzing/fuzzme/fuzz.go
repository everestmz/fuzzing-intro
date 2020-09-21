package fuzzme

func Fuzz(Data []byte) int {
	BrokenMethod(string(Data))
	return 0
}
