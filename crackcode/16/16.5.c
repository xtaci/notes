#include <stdio.h>

#define BIG_ENDIAN 0
#define LITTLE_ENDIAN 0

int byteorder();

int
main(void) {
	printf("%d\n", byteorder());
}

int
byteorder() {
	short word = 0x0001;
	char * byte = (char *) &word;
	return (byte[0] ? LITTLE_ENDIAN:BIG_ENDIAN);
}
