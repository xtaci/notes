#include <stdio.h>
#include <string.h>

const int N = 10;
enum Color {
	Green = 1
};

void fill(int screen[][N], int x, int y, enum Color color);
void print_m(int m[][N]);

int 
main(void) {
	int screen[N][N];
	memset(screen ,0, sizeof(screen));
	printf("sizeof 2d-array %ld\n", sizeof(screen));
	print_m(screen);
	fill(screen ,N-1, N-1, Green);
	print_m(screen);
}

void
fill(int screen[][N], int x, int y, enum Color color) {
	if (x <0 || x > N || y <0 || y > N) {
		return;
	}

	if (screen[y][x] == 0) {
		screen[y][x] = color;
		fill(screen, x-1, y, color);
		fill(screen, x+1, y, color);
		fill(screen, x, y-1, color);
		fill(screen, x, y+1, color);
	}
}

void
print_m(int m[][N]) {
	for(int y = 0;y<N;y++) {
		for(int x = 0;x<N;x++) {
			printf("%d ", m[y][x]);
		}
		printf("\n");
	}
}
