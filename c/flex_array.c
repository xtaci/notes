#include <stdio.h>
#include <stdlib.h>

struct packet {
	unsigned short len;
	char p[];
};

int 
main(void) {
	printf("%ld\n", sizeof(struct packet));
	struct packet * s = (struct packet *)malloc(sizeof(struct packet) + 10);
	s->len = 10;
	for (int i=0;i<s->len;i++) {
		printf("%hhd ", s->p[i]);
	}
}
