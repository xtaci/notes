package main

const (
	FRAGMENT_SIZE = 256
)

//------------------------------------------------ 分割大的tcp包
func fragment(data []byte) (segments [][]byte) {
	sz := len(data)
	quotient := sz / FRAGMENT_SIZE
	remainder := sz % FRAGMENT_SIZE

	for i := 0; i < quotient; i++ {
		segments = append(segments, data[i*FRAGMENT_SIZE:(i+1)*FRAGMENT_SIZE])
	}

	if remainder > 0 {
		segments = append(segments, data[quotient*FRAGMENT_SIZE:])
	}
	return
}
