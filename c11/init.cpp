#include <iostream>

class T {
	public:
	int A;
	int B;
	T();
};

T::T() {
}

int 
main() {
	using namespace std;
	T *x = new(T);
	T t;
	int a;
	cout << x->A << " " << x->B  << endl;
	cout << t.A << " " << t.B  << endl;
	cout << a << endl;
	int a3[5] = {0,1,2}; 
	int b3[5];
	for (auto &v :a3) {
		cout << v << " ";
	}
	cout << endl;	
	for (auto &v :b3) {
		cout << v << " " ;
	}
	cout << endl;	
}
