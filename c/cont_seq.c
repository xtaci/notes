#include <stdio.h>

void PrintContinuousSequence(int small, int big);
void FindContinuousSequence(int n);

int 
main() {
	FindContinuousSequence(15);
}

void FindContinuousSequence(int n) {
	if(n < 3)
		return;

	int small = 1; 
	int big = 2;
	int sum = small + big;

	while(small < big) {
		// we are lucky and find the sequence
		if(sum == n)
			PrintContinuousSequence(small, big);

		// if the current sum is greater than n, 
		// move small forward
		while(sum > n) {
			sum -= small;
			small ++;

			// we are lucky and find the sequence
			if(sum == n)
				PrintContinuousSequence(small, big);
		}

		// move big forward
		big ++;
		sum += big;
	}
}

/////////////////////////////////////////////////////////////////////////
// Print continuous sequence between small and big
/////////////////////////////////////////////////////////////////////////
void PrintContinuousSequence(int small, int big) {
	for(int i = small; i <= big; ++ i)
		printf("%d ", i);

	printf("\n");
}
