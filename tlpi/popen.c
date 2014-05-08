#include <stdio.h>

int main(void) {
	char buf[1024];
	FILE * fp = popen("/bin/ls -la", "r");
	if (fp == 0) {
		perror("popen");
	}

	while(!feof(fp)) {
		size_t n = fread(buf, 1, 1024, fp);
		write(1, buf, n);
	}
	pclose(fp);
}
