#include "sqrt_test.h"

double norm2(double x, double y) {
  double n1 = x * x + y * y;
  double n = (double)sqrt(n1);
  return n;
}
int norm2_int(double x, double y) { return (int)sqrt(x * x + y * y); }
float norm2_fl(double x, double y) { return (float)sqrt(x * x + y * y); }
int sz(size_t var) { return sizeof(var); }