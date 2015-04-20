#include <stdio.h>

void merge(int a[], int b[], int n, int m);

int
main(void) {
	int a[10] = {2,4,6,8,10};
	int b[] = {1,3,5,7,9};
	merge(a,b,5,5); 
	for (int i=0;i<sizeof(a)/sizeof(int);i++) {
		printf("%d ", a[i]);
	}
}

void 
merge(int a[], int b[], int n, int m) {
	int k = m+n-1;
	int i = n-1;
	int j = m-1;

	while(i>=0 && j>=0) {
		if (a[i] > b[j]) {
			a[k--] = a[i--];
		} else {
			a[k--] = b[j--];
		}
	}

	while (j>=0) {
		a[k--] = b[j--];
	}
}
