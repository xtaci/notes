#include <stdio.h>

const int N = 4;

void print_m(const int m[][N]);
void rotate_m(int m[][N]);

int
main(void) {
	int m[N][N];
	for (int i=0;i<N;i++) {
		for (int j=0;j<N;j++) {
			m[i][j] = i*j;
		}
	}
	print_m(m);
	rotate_m(m);
	print_m(m);
}

void
print_m(const int m[][N]) {
	for (int i=0;i<N;i++) {
		printf("%d %d %d %d\n", m[i][0], m[i][1],m[i][2], m[i][3]);
	}
}


void
rotate_m(int m[][N]) {
	for (int i=0;i<N;i++) {
		for (int j=0;j<N-i;j++) {
			int b = m[i][j];
			m[i][j] = m[N-j-1][N-i-1];
			m[N-j-1][N-i-1] = b;
		}
	}
}
