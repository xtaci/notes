#include <stdio.h>

int bitdiff(int a, int b);

int
main(void) {
	int a = 31;
	int b = 14;
	printf("bitdiff 31, 14 is %d\n", bitdiff(a,b));
}

int 
bitdiff(int a, int b) {
	int c = a ^b;
	int count = 0;
	while(c!=0) {
		count += c&1;
		c >>= 1;
	}
	return count;
}
