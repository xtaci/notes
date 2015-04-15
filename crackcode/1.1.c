#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <stdbool.h>

bool unique(char * s, size_t len);

int 
main(void) {
	char s1[] = "hello";
	char s2[] = "thanks";
	printf("%d\n", unique(s1, sizeof(s1)-1));
	printf("%d\n", unique(s2, sizeof(s2)-1));
}

bool 
unique(char *restrict s, size_t len) {
	bool test[256];
	memset(test, 0, sizeof(bool) * 256);

	for (int i=0;i<len;i++) {
		if (test[s[i]]) {
			return false;
		}

		test[s[i]] = true;
	}
	return true;	
}
