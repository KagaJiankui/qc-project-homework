#pragma once
#define _CRT_SECURE_NO_WARNINGS
#ifndef __SQRT_TEST_H
#define __SQRT_TEST_H
#include <cmath>
#ifdef __cplusplus
extern "C" {
#endif  // __cplusplus
__declspec(dllexport) double norm2(double x, double y);
__declspec(dllexport) int norm2_int(double x, double y);
__declspec(dllexport) float norm2_fl(double x,
                                     double y);  //测试不同长度的返回值
__declspec(dllexport) int sz(size_t var);
#ifdef __cplusplus
}
#endif  // __cplusplus
#endif
