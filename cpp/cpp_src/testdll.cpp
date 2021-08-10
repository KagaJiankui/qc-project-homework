#include <Windows.h>
#include <stdio.h>
// #include <tchar.h>

#include "sqrt_test.h"
using namespace std;
typedef double (*pNorm)(double, double);

int main(int argc, char const* argv[]) {
  const char* dllName = "DLLTest1.dll";
  const char* procName = "norm2";
  HMODULE hDLL = LoadLibrary(dllName);
  if (hDLL != NULL) {
    pNorm fp = (pNorm)GetProcAddress(hDLL, procName);
    if (fp == NULL) {
      printf("Cannot Find Required Function\n");
    } else {
      double a = 3.3;
      double b = 1.4;
      printf("result: %.2f", fp(a, b));
    }
    FreeLibrary(hDLL);
    return 0;
  } else {
    return 1;
  }
}
