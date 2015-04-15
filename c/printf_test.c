#include <stdio.h>
#include <stddef.h>

int 
main(void) {
	printf("%2$d, %1$d\n", 1, 2);	// specify positions of arguments
	int count;
	printf("the quick brown fox jumps over the lazy dog. %n", &count);
	printf("count: %d\n",count);
	printf("%+d %+d\n", 10, -10);
	printf("%03d\n", 10);
	float pi = 3.141592653584763;
	printf("%.f %.2f %0.1f \n", pi, pi, pi);

	char a = 'A';
	short b = 'B';
	int c = 'C';
	long d = 'D';
	long long e = 'E';
	printf("%hhd %hd %d %ld %lld\n", a,b,c,d,e);

	printf("convert %f to %6e, print with g %g %g\n", pi, pi*100, 123456789.0f, 0.0000001);
	printf("convert with a: %a\n", pi);
	printf("print pointer %p\n", main);

	long double ld = pi;
	printf("long double: %Lf\n", ld);

	ptrdiff_t diff = (void*)printf- (void*)main;
	printf("%p, %p, diff : %lx\n", main, printf, diff);
}
