#include <iostream>

class test{
	int a{0};
	test();
};

test::test():a(1){
}

int
main(void) {
	auto t = new(test);
}
