#include <stdio.h>

int 
main(int argc, char * argv[]) {
	extern char etext, edata, end;
	printf("%x %x %x\n", &etext, &edata, &end);
}
