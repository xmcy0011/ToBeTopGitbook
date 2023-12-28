cmake_minimum_required(VERSION 2.8)
project(04_json)

set(CMAKE_CXX_STANDARD 14)
add_compile_options(-std=c++11)

include_directories(.)
aux_source_directory(. SRC)

add_executable(04_json main.cpp ${SRC})