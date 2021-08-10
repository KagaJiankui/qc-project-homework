cmake_minimum_required(VERSION 2.6)
project(DLLTest1)

aux_source_directory("./cpp_src" SRC_LIST)
set(TEST_SOURCE "./cpp_src/testdll.cpp")
include_directories("./cpp_incl")

add_executable(DLLTest1a ${TEST_SOURCE})
add_library(DLLTest1 SHARED ${SRC_LIST})