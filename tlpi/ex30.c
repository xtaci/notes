#include <stdio.h>

int main(void) 
{ 
	int x ;
    printf("%d %d %d", x, 1&&x, (x == (1&&x)));
	return 0;
}
