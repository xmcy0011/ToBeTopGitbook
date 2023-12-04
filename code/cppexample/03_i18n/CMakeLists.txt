cmake_minimum_required(VERSION 2.8)
project(01_share_ptr_oop)

set(CMAKE_CXX_STANDARD 14)
add_compile_options(-std=c++11)

include_directories(.)
aux_source_directory(. SRC)

add_executable(02_toml main.cpp ${SRC})