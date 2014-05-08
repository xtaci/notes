#include <stdio.h>

int main(void) {
	char buf[1024];
	FILE * fp = popen("/bin/ls -la", "r");
	if (fp == 0) {
		perror("popen");
	}

	while(fgets(buf, 1024, fp)!=0) {
		printf("%s",buf);
	}
	pclose(fp);
}
