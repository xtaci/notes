#include <stdio.h>
#include <string.h>

char * removedup(char * restrict s, int n);

int
main(void) {
	char s1[] = "hello";
	char s2[] = "well";
	char s3[] = "this is a sentence";
	printf("%s\n", removedup(s1, sizeof(s1)-1));
	printf("%s\n", removedup(s2, sizeof(s2)-1));
	printf("%s\n", removedup(s3, sizeof(s3)-1));
}

char * removedup(char * restrict s, int n) {
	for (int i=0;i<n;i++) {
		for (int j=i+1;j<n;j++) {
			if (s[i] == s[j]) {
				memmove(s+j,s+j+1,n-j-1);
				j--;
				n--;
			}
		}
	}
	s[n] = 0;
	return s;
}
