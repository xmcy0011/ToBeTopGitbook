#include <iostream>

int main() {
  int dirType = 1;
  int fileType = dirType << 1;
  int dirOrFileType = dirType | fileType;
  
  int objType = 0;
  bool val = !(objType & dirOrFileType);

  std::cout << "dirType:" << dirType << ",fileType:" << fileType << ",dirOrFileType:" << dirOrFileType << std::endl;
  std::cout << "objType is file or dir ?" << val << std::endl;
}
