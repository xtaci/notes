#include <stdio.h>

char * reverse(char * restrict s, size_t n);

int 
main(void) {
	char s1[] = "bananas";
	char s2[] = "temp";
	printf("%s\n", reverse(s1, sizeof(s1)-1));
	printf("%s\n", reverse(s2, sizeof(s2)-1));

	printf("%s\n", reverse(s1, sizeof(s1)-1));
	printf("%s\n", reverse(s2, sizeof(s2)-1));
}

char *
reverse(char * restrict s, size_t n) {
	for (int i=0;i<n/2;i++) {
		char b = s[i];
		s[i] = s[n-i-1];
		s[n-i-1] = b;
	}
	return s;
}
