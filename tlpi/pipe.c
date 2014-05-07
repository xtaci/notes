#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <sys/time.h>

#define BUFSIZE 4096

int main(void) {
	int filedes[2];

	char * readbuf = malloc(BUFSIZE);
	if (readbuf== NULL) {
		perror(0);
		exit(-1);
	}

	char * dummy = malloc(BUFSIZE);
	if (dummy== NULL) {
		perror(0);
		exit(-1);
	}

	if (pipe(filedes) == -1) {
		perror(0);	
		exit(-1);
	}

	pid_t pid = fork();

	struct timeval begin;
	struct timeval now;
	gettimeofday(&begin, 0);

	int64_t count = 0;
	if (pid == 0) {
		while(1) {
			int n = read(filedes[0], readbuf, BUFSIZE);
			count +=n;
			gettimeofday(&now, 0);

			int interval = now.tv_sec - begin.tv_sec;
			if (interval > 1) {
				printf("%f MB\r", (float)count/(interval*1024*1024));
			}
		}
	} else {
		while(1) {
			write(filedes[1], dummy, BUFSIZE);
		}
	}
	return 0;
}
