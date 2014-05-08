#include <stdio.h>
#include <fcntl.h>
#include <unistd.h>

int main(void) {
	for (;;) {
		int fd = open("/tmp/test", O_RDWR|O_CREAT);
		if (fd == -1) {
			perror("open");
			return -1;
		}

		printf("fd := %d\n", fd);
		if (dup2(1,fd) == -1) {
			perror("dup2");
		}
		close(fd);
	}

	return 0;
}
