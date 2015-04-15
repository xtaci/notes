#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>

void my_memset(void * s, char c,int n);

int
main(void) {
	const int N = 199;
	char *s = (char *)malloc(N);
	my_memset(s, 0xAA, N);
	for (int i=0;i<N;i++) {
		printf("%hhx ", s[i]);
	}
}

void
my_memset(void * s, char c, int n) {
	int i;
	int ws = sizeof(uintptr_t);
	uintptr_t expand = 0;
	for (i=0;i<ws;i++) {
		expand |= (0xff&(uintptr_t)c)<<8*i;
	}

	for (i=0;i<n/ws;i++) {
		*(uintptr_t*)(s+ws*i) = expand;
	}

	for (i=n%ws;i>0;i--) {
		*(char*)(s+n-i) = c;
	}
}
