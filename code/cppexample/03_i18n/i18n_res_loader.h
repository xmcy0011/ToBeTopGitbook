#ifndef CODE_CPPEXAMPLE_03_I18N_I18N_RES_LOADER_H_
#define CODE_CPPEXAMPLE_03_I18N_I18N_RES_LOADER_H_

namespace cpptoml {
class table;
};

#include <memory>
#include <string>
#include <unordered_map>

/**
 * @brief 多语言管理器
 */
class I18NResLoader {
  private:
    I18NResLoader() = default;
    ~I18NResLoader() = default;

  public:
    static I18NResLoader &GetInstance();

  public:
    /**
     * @brief 加载多语言资源，如果失败直接报错
     * @param lang         : 语言
     * @param filePath     : 资源 toml 文件路径
     */
    void Load(const std::string &lang, const std::string & filePath);

    /**
     * @brief 加载资源
     * @param lang         : 语言
     * @param key          : 资源 key
     * @return 对应语言的 value 值，如果没找到返回空字符串
     */
    std::string Localize(const std::string & lang, const std::string & key) noexcept;

  private:
    std::unordered_map<std::string, std::shared_ptr<cpptoml::table>> res_;
};

#endif // CODE_CPPEXAMPLE_03_I18N_I18N_RES_LOADER_H_