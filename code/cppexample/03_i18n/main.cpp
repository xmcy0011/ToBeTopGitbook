#include "cpptoml.h"

#include <iostream>
#include <memory>
#include <string>

#include "http_rest_exception.h"
#include "i18n_res_loader.h"

const std::string kSomeError = "efast.err.in.approval";

void testConflictExecption() { throw ConflictException(__FILE__, __LINE__, "eofs", kSomeError); }

int main(int argc, char **argv) {

    std::string enUSLan = "en-US";
    std::string zhCNLan = "zh-CN";

    std::string enUSFile = "../locale/example-en-US.toml";
    std::string zhCNFile = "../locale/example-zh-CN.toml";

    auto &localizer = I18NResLoader::GetInstance();
    localizer.Load(enUSLan, enUSFile);
    localizer.Load(zhCNLan, zhCNFile);

    std::cout << "en-US: " << localizer.Localize(enUSLan, kSomeError) << std::endl;
    std::cout << "zn-CH: " << localizer.Localize(zhCNLan, kSomeError) << std::endl;

    std::string resourceNotFoundKey = "efast.err.in.approvalxxxxxxxxxxxx";
    std::cout << "en-US: " << localizer.Localize(enUSLan, resourceNotFoundKey) << std::endl;
    std::cout << "zn-CH: " << localizer.Localize(zhCNLan, resourceNotFoundKey) << std::endl;

    try {
        testConflictExecption();
    } catch (ClientException e) {
        std::cout << "client execption: " << localizer.Localize(enUSLan, e.GetResouceKey()) << std::endl;
    } catch (ServerException e) {
        std::cout << "client execption: " << localizer.Localize(enUSLan, e.GetResouceKey()) << std::endl;
    }

    return 0;
}
