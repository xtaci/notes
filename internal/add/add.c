#include "runtime.h"

void ·Add(uint64 a, uint64 b, uint64 ret) {
	ret = a + b;
	FLUSH(&ret);
}
