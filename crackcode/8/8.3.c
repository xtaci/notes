#include <stdio.h>
#include <string.h>

void subset(char s[], char onoff[], size_t n);

int 
main(void) {
	char s[] = "abcd";
	char onoff[sizeof(s)-1];
	subset(s, onoff, sizeof(s)-1);
}

void 
subset(char s[], char onoff[], size_t n) {
	if (n == 0) {
		for (int i = 0;i<strlen(s);i++) {
			if (onoff[i] == 1) {
				printf("%c", s[i]);
			}
		}
		printf("\n");
		return;
	}

	onoff[strlen(s)-n] = 0;
	subset(s, onoff, n-1);

	onoff[strlen(s)-n] = 1;
	subset(s, onoff, n-1);
}
