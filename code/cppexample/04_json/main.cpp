#include <iostream>

#include "json.hpp"
using json = nlohmann::json;

int main() {
    json j;
    j["name"] = "第八届全国脊柱内镜大会暨中国医疗保健国际交流促进会骨科分会脊柱内镜学部2023年会暨河南省康复医学会骨科微创专委会成立大会";
    j["age"] = 12;

    json a = json::array();
    a.emplace_back(1);

    j["arr"] = json::array();
    j["arr2"] = a;

    std::cout << j.dump() << std::endl;

    return 0;
}