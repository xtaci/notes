#include <stdio.h>

void print_bin(const char * decimal, size_t n);
void print_bin2(int n);

int 
main(void) {
	char s[] = "3.72";
	print_bin(s, sizeof(s)-1);
}

void
print_bin(const char * decimal, size_t n) {
	int i;
	int left = 0;
	int right = 0;

	for (i=0;i<n;i++) {
		if (decimal[i] == '.') {
			break;
		}
		left *= 10;	
		left += decimal[i] - '0';
	}
	i++;	// skip dot

	for (;i<n;i++) {
		if (decimal[i] == '.') {
			break;
		}
		right *= 10;	
		right += decimal[i] - '0';
	}

	print_bin2(left);
	printf(".");
	print_bin2(right);
}

void 
print_bin2(int n) {
	if (n > 0) {
		print_bin2(n>>1);
	}

	if (n!=0) {
		printf("%d", n&1);
	}
}
