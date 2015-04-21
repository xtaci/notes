#include <stdio.h>
#include <stdlib.h>

int ** alloc2d(int rows, int cols);

int
main(void) {
	int **ptr = alloc2d(10,10);
	for (int i=0;i<10;i++) {
		for (int j=0;j<10;j++) {
			printf("%d ", ptr[i][j]);
		}
		printf("\n");
	}
}

int ** alloc2d(int rows, int cols) {
	int header = rows * sizeof(int*);
	int data = rows * cols * sizeof(int);
	int **rowptr =  (int **)malloc(header+data);
	int *buf = (int*)(rowptr+rows);
	int k;
	for (k=0;k<rows;k++) {
		rowptr[k] = buf + k*cols;
	}
	return rowptr;
}
