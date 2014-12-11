package main

func alloc_heap() []byte {
	ah := make([]byte, 10)
	ah[1] = 1
	return nil
}

func alloc_stack(v []byte) {
	v[1] = 1
}

func main() {
	alloc_heap()

	s := make([]byte, 10)
	alloc_stack(s)
}
