#include <stdio.h>

void print_par(int l, int r, char str[], int count);

int 
main(void) {
	const int N = 3;
	char s[N*2+1];
	s[N*2] = 0;
	print_par(N, N, s, 0);
}

void 
print_par(int l, int r, char str[], int count) {
	if (l<0 || r <l) return;
	if (l == 0 && r == 0) {
		printf("%s     ", str);
		return;
	}

	if (l>0) {
		str[count] = '(';
		print_par(l-1, r, str, count+1);
	}

	if (r>l) {
		str[count] = ')';
		print_par(l, r-1, str, count+1);
	}
}
