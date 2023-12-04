#include "cpptoml.h"

#include <iostream>
#include <memory>
#include <string>
#include <unordered_map>

// class I18NLoader {
//   public:
//     static const I18NLoader& getInstance();

//   private:
//     I18NLoader() = default;
//     ~I18NLoader() = default;
// };

int main(int argc, char **argv) {

    std::string file = "../example.toml";
    std::shared_ptr<cpptoml::table> filePtr = nullptr;

    try {
        filePtr = cpptoml::parse_file(file);
        std::cout << (*filePtr) << std::endl;
    } catch (const cpptoml::parse_exception &e) {
        std::cerr << "Failed to parse " << argv[1] << ": " << e.what() << std::endl;
        return 1;
    }

    std::string resourceKey = "efast.err.in.approval";
    std::cout << "key1: " << filePtr->get_qualified_as<std::string>(resourceKey)->c_str() << std::endl;

    std::string resourceNotFoundKey = "efast.err.in.approvalxxxxxxxxxxxx";
    std::cout << "key2: " << filePtr->get_qualified_as<std::string>(resourceNotFoundKey)->c_str() << std::endl;

    return 0;
}
