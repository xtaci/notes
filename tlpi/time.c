#include <time.h>
#include <stdlib.h>
#include <stdio.h>

int main(void) {
	time_t x = time(NULL);
	printf("%s\n", ctime(&x));
	struct tm * t = localtime(&x);
	printf("%d %d %d %d %s\n", t->tm_sec, t->tm_min, t->tm_hour, t->tm_mday, t->tm_zone);
	t = gmtime(&x);
	printf("%d %d %d %d %s\n", t->tm_sec, t->tm_min, t->tm_hour, t->tm_mday, t->tm_zone);
	exit(0);
}
