#include "i18n_res_loader.h"
#include "cpptoml.h"
#include <dirent.h>
#include <exception>

I18NResLoader &I18NResLoader::GetInstance() {
    static I18NResLoader loader;
    return loader;
}

void I18NResLoader::LoadFile(const std::string &lang, const std::string &filePath) {
    auto ptr = cpptoml::parse_file(filePath);
    res_[lang] = ptr;
}

void GetFileNames(std::string path, std::vector<std::string> &filenames) {
    DIR *pDir;
    struct dirent *ptr;
    if (!(pDir = opendir(path.c_str()))) {
        throw std::invalid_argument("Folder doesn't Exist!");
    }
    while ((ptr = readdir(pDir)) != 0) {
        if (strcmp(ptr->d_name, ".") != 0 && strcmp(ptr->d_name, "..") != 0) {
            filenames.push_back(path + "/" + ptr->d_name);
        }
    }
    closedir(pDir);
}

void I18NResLoader::LoadDir(const std::string &langDir) {
    if (loaded_) {
        return;
    }

    std::vector<std::string> files;
    GetFileNames(langDir, files);
    for (auto &&f : files) {
        if (f.find("toml") == std::string::npos) {
            continue;
        }

        // en-US.toml -> en-US
        size_t pos = f.find_first_of('.');
        std::string lan = f;
        if (pos != std::string::npos) {
            lan = f.substr(0, pos);
        }

        // load
        LoadFile(lan, f);
    }

    loaded_ = true;
}

std::string I18NResLoader::Localize(const std::string &lang, const std::string &key) noexcept {
    std::string text;
    if (res_.find(lang) != res_.end()) {
        text = res_[lang]->get_qualified_as<std::string>(key)->c_str();
    }
    return std::move(text);
}