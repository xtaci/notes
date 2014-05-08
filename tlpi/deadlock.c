#include <stdio.h>
#include <signal.h>
#include <pthread.h>
#include <stdint.h>

static pthread_mutex_t mtx = PTHREAD_MUTEX_INITIALIZER;

void sig_handler(int);

int main(void) {
	// init signal
	signal(SIGUSR1, sig_handler);

	// lock
	pthread_mutex_lock(&mtx);
	int64_t i = 0;
	for (;;) {
		printf("%ld\r", i++);
	}	
	pthread_mutex_unlock(&mtx);
}

void sig_handler(int sig) {
	printf("signal handler entered %d\n",sig);
	pthread_mutex_lock(&mtx);
	printf("signal handler mutex entered %d\n", sig);
	pthread_mutex_unlock(&mtx);
}
