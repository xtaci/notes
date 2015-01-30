#include <vector>
#include <iostream>

using namespace std;
class Solution {
	public:
		vector<int> twoSum(vector<int> &numbers, int target) {
			int sz = numbers.size();
			vector<int> result;
			for (int i=0;i<sz;i++) {
				int v = target - numbers[i];
				for (int j=i+1;j<sz;j++) {
					if (numbers[j] == v) {
						result.push_back(i+1);
						result.push_back(j+1);
						return result;
					}
				}
			}

			return result;
		}
};

int main(void) {
	vector<int> numbers;
	numbers.push_back(2);
	numbers.push_back(7);
	numbers.push_back(11);
	numbers.push_back(15);
	Solution s;
	vector<int> ret =  s.twoSum(numbers, 9);
	for (vector<int>::iterator iter = ret.begin();iter != ret.end();++iter) {
		cout << *iter << endl;
	}

	return 0;
}
