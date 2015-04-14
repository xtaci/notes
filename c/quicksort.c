#include <stdio.h>
#include <stdlib.h>

void quicksort(int list[], int begin, int end);
int partition(int list[], int begin, int end);
void swap(int *x, int *y);

int 
main(void) {
	int test[] = {5,1,3,9,2,4,10,8};
	int length = sizeof(test)/sizeof(int);
	quicksort(test, 0, length-1);
	for (int i=0;i<length;i++) {
		printf("%d ", test[i]);
	}
	printf("\n");
}

void
swap(int *x, int *y) {
	int tmp = *x;
	*x = *y;
	*y = tmp;
}

void 
quicksort(int list[], int begin, int end) {
		if (begin < end) {
			int pivot_idx = partition(list, begin, end);
			quicksort(list, begin, pivot_idx -1);
			quicksort(list, pivot_idx +1, end);
		}
}

int 
partition(int list[], int begin, int end) {
	int pivot_idx = rand()%(end-begin+1) + begin;
	int pivot = list[pivot_idx];
	// swap
	swap(&list[pivot_idx], &list[begin]);

	int i = begin +1;
	int j = end;

	while(i<=j) {
		while( (i <= end) && (list[i] <= pivot)) i++;
		while( (j >= begin) && (list[j] > pivot)) j--;
		if (i < j) swap(&list[i], &list[j]);
	}

	swap(&list[begin], &list[j]);
	return j;
}
