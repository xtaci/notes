#include <stdio.h>
#include <stdint.h>

int 
main(void) {
	printf("sizeof int %ld\n", sizeof(int));
	printf("sizeof long %ld\n", sizeof(long));
	printf("sizeof long long %ld\n", sizeof(long long));
	printf("sizeof uintptr_t %ld\n", sizeof(uintptr_t));
	printf("sizeof void * %ld\n", sizeof(void*));
}
