#include <stdio.h>
#include <stdint.h>

uint32_t setbits(uint32_t n, uint32_t m, int i,int j);

int
main(void) {
	printf("%d", setbits(1024,21,2,6));
}

uint32_t
setbits(uint32_t n, uint32_t m, int i,int j) {
	uint32_t mask =  ~0;
	mask = ~((mask >> (32-(j-i+1))) << i);

	n &= mask;
	n |= (m<<i) &~mask;
	return n;
}
