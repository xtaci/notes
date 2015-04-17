#include <stdio.h>
#include <stdint.h>

uint32_t swapbits(uint32_t n);

int 
main(void) {
	uint32_t n = 1234;
	printf("swap odd even bits %x into %x",  n, swapbits(n));
}

uint32_t 
swapbits(uint32_t n) {
	uint32_t mask1 = 0xaaaaaaaa;
	uint32_t mask2 = 0x55555555;

	return ((n&mask1) >> 1) | ((n&mask2) <<1);
}
