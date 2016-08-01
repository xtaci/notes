#include <iostream>

struct T {
	int A;
};

int 
main() {
	T *x = new(T);
	int a;
	std::cout << x->A << std::endl;
	std::cout << a << std::endl;
	int a3[5] = {0,1,2}; 
	for (auto &v :a3) {
		std::cout << v <<std::endl;
	}
}
