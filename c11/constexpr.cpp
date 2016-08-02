#include <iostream>

inline const int
get_size() {
	return 10;
}

int
main(void) {
	unsigned sz=42 ;
	constexpr unsigned sz1 = 42;
	const unsigned sz2 = 42;
	char a[sz];
	char b[sz1];
	char c[sz2];
	char d[get_size()];

	std::string str1[sz];
	std::string str2[sz1];
	std::string str3[sz2];
	std::string str4[get_size()];
}
