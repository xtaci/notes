#include <stdio.h>
#include <stdbool.h>

bool anagram(char * restrict s1, int n1, char * restrict s2, int n2);

int
main(void) {
	char s1[] = "mary";
	char s2[] = "army";
	char s3[] = "mami";
	printf("%d\n", anagram(s1, sizeof(s1)-1, s2, sizeof(s2)-1));
	printf("%d\n", anagram(s1, sizeof(s1)-1, s3, sizeof(s3)-1));
}

bool
anagram(char * restrict s1, int n1, char * restrict s2, int n2) {
	if (n1!=n2) {
		return false;
	}

	for (int i=0;i<n1;i++){
		for (int j=0;j<n2;j++) {
			if (s1[i] == s2[j]) {
				s2[j] = 0;
			}
		}
	}

	for (int j=0;j<n2;j++) {
		if (s2[j] > 0) return false;
	}
	return true;
}
