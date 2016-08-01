#include <iostream>

const int
get_size() {
	return 10;
}

int
main(void) {
	char d[get_size()];
	const int cnt = 42;
	std::string x[cnt];
	std::string y[get_size()];
}

