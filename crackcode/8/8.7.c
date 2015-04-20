#include <stdio.h>

int make_change(int n, int denom);

int 
main(void) {
	printf("%d\n", make_change(100, 25));
}


int
make_change(int n, int denom) {
	int next_denom = 0;
	switch (denom) {
	case 25:
		next_denom = 10;
		break;
	case 10:
		next_denom = 5;
		break;
	case 5:
		next_denom = 1;
		break;
	case 1:
		return 1;
	}

	int ways = 0;
	for (int i=0;i*denom <=n;i++) {
		ways += make_change(n - i*denom, next_denom);
	}
	return ways;
}
