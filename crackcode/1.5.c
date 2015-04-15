#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char * quote(char * restrict s, size_t n);

int 
main(void) {
	char s1[] = "this is a sentence";
	char *q = quote(s1, sizeof(s1)-1);
	printf("%s %ld\n", q, strlen(q));
}

char * 
quote(char * restrict s, size_t n) {
	int num_spaces = 0;
	for (int i=0;i<n;i++) {
		if (s[i] == ' ') {
			num_spaces++;
		}
	}

	char * quote_s = (char *) malloc(n+num_spaces*2+1);
	int cnt = 0;
	for (int i=0;i<n;i++) {
		if (s[i] != ' ') {
			quote_s[i+cnt*2] = s[i];
		} else {
			int base = cnt *2 +i;
			quote_s[base] = '%';
			quote_s[base+1] = '2';
			quote_s[base+2] = '0';
			cnt++;
		}
	}

	quote_s[n+num_spaces*2] = 0;
	return quote_s;
}
