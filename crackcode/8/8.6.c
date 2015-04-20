#include <stdio.h>
#include <string.h>

const int N = 10;
enum Color {
	Black = 0,
	Green = 1
};

void fill(int screen[][N], int x, int y, enum Color ocolor, enum Color ncolor);
void print_m(int m[][N]);

int 
main(void) {
	int screen[N][N];
	memset(screen ,Black, sizeof(screen));
	printf("sizeof 2d-array %ld\n", sizeof(screen));
	print_m(screen);
	fill(screen ,N-1, N-1, screen[N-1][N-1], Green);
	print_m(screen);
}

void 
fill(int screen[][N], int x, int y, enum Color ocolor, enum Color ncolor) {
	if (x <0 || x > N || y <0 || y > N) {
		return;
	}

	if (screen[y][x] == ocolor) {
		screen[y][x] = ncolor;
		fill(screen, x-1, y, ocolor, ncolor);
		fill(screen, x+1, y, ocolor, ncolor);
		fill(screen, x, y-1, ocolor, ncolor);
		fill(screen, x, y+1, ocolor, ncolor);
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
