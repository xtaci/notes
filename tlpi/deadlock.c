#include <stdio.h>
#include <signal.h>
#include <stdint.h>
#include <pthread.h>


void sig_handler(int);
void *worker(void *);

int main(void) {
	// init signal handler
	signal(SIGHUP, sig_handler);

	pthread_t t1, t2;
	pthread_create(&t1, 0, worker, 0);
	pthread_create(&t2, 0, worker, 0);

	int64_t i = 0;
	for (;;) {
		printf("%ld\r", i++);
	}	
}

void sig_handler(int sig) {
	for (;;) {
	}
}

void *worker(void * arg) {
	int64_t i = 0;
	for (;;) {
		printf("thread %d\t%ld\r", pthread_self(), i++);
	}	
}
