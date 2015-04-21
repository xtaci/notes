#include <stdio.h>

int 
main(void) {
	int s[][2] = {{1,2},{3,4},{5,6}};
	for (int i =0;i<3;i++) {
		for (int j=0;j<2;j++) {
			printf("%d %p    ", s[i][j], &s[i][j]);
		}
		printf("\n");
	}
}
