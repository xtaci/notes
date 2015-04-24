#include <stdio.h>

int max_subarray(int s[], size_t n);

int 
main(void) {
	int s[] = {1, -2, 3, 10, -4, 7, 2, -5};
	printf("max: %d\n", max_subarray(s, sizeof(s)/sizeof(int)));
}

int
max_subarray(int s[], size_t n) {
	int result = s[0];
	int sum = s[0];
	for (int i =1;i<n;i++) {
		if (sum > 0) {
			sum += s[i];
		} else {
			sum = s[i];
		}

		if (sum > result) {
			result =sum;
		}
	}
	return result;
}
