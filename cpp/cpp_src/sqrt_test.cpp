#include "sqrt_test.h"

double norm2(double x, double y) {
  double n1 = x * x + y * y;
  double n = (double)sqrt(n1);
  return n;
}