#include <stdio.h>
#include <string.h>
#include <stdbool.h>
#include <stdlib.h>

bool is_rotate_str(const char * s1,size_t n1, const char * s2, size_t n2);
bool isSubString(const char *s1, const char *s2);

int
main(void) {
	char s1[] = "erbottlewat";
	char s2[] = "waterbottle";
	char s3[] = "waterbottlee";

	printf("%d\n", is_rotate_str(s1, sizeof(s1)-1, s2, sizeof(s2)-1));
	printf("%d\n", is_rotate_str(s1, sizeof(s1)-1, s3, sizeof(s3)-1));
}

bool 
is_rotate_str(const char * s1,size_t n1, const char * s2, size_t n2) {
	char * s = (char *)malloc(n1*2+1);
	for(int i=0;i<n1*2;i++) {
		s[i] = s1[i%n1];
	}
	s[n1*2] = 0;

	return isSubString(s, s2);
}

bool
isSubString(const char *s1, const char *s2) {
	return strstr(s1,s2)>0? true: false;
}
