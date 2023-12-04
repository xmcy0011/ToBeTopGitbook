#include "i18n_res_loader.h"
#include "cpptoml.h"

I18NResLoader &I18NResLoader::GetInstance() {
    static I18NResLoader loader;
    return loader;
}

void I18NResLoader::Load(const std::string &lang, const std::string &filePath) {
    auto ptr = cpptoml::parse_file(filePath);
    res_[lang] = ptr;
}

std::string I18NResLoader::Localize(const std::string &lang, const std::string &key) noexcept {
    std::string text;
    if (res_.find(lang) != res_.end()) {
        text = res_[lang]->get_qualified_as<std::string>(key)->c_str();
    }
    return std::move(text);
}