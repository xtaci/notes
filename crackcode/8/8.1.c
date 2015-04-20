#include <stdio.h>

long fib(long n);

int 
main(void) {
	printf("%ld\n", fib(10));
}

long
fib(long n) {
	if (n<0) return -1;
	if (n==0) return 0;
	long a = 1, b=1;
	for (long i=3;i<=n;i++) {
		long c = a+b;
		a=b;
		b=c;
	}
	return b;
}
