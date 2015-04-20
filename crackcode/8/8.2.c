#include <stdio.h>

long robot(int x, int y);

int 
main(void) {
	const N = 4;

	printf("%ld\n", robot(N,N));
}

long
robot(int x, int y) {
	if (x ==0 && y == 0) {
		return 0;
	}

	long count = 0;
	if (y > 0) count = 1 + robot(x,y-1);
	if (x > 0) count += 1+ robot(x-1, y);
	return count;
}
