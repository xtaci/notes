#include <stdio.h>

void print_p(char s[], size_t sz, int depth);
void swap(char s[], int i, int j);

int
main(void) {
	char s[] = "abc";
	print_p(s, sizeof(s)-1, 0);
}

void 
print_p(char s[], size_t sz, int depth) {
	if (depth==sz-1) {
		printf("%s\n", s);
		return;
	}

	print_p(s, sz, depth+1);

	for (int i = depth+1;i<sz;i++) {
		swap(s,depth,i);
		print_p(s, sz, depth+1);
		swap(s,depth,i);
	}

	return;
}

void
swap(char s[], int i, int j) {
	int tmp;
	tmp = s[i];
	s[i] = s[j];
	s[j] = tmp;
}
