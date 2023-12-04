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
    void LoadFile(const std::string &lang, const std::string & filePath);

    /**
     * @brief 加载多语言资源
     * @param filePath     : 资源 toml 文件夹路径
     * @example 文件夹下资源必须符合如下规则：
     * [lan].toml 即文件名为语言名称
     * 如：
     * en-US.toml
     * zh-CN.toml
     */
    void LoadDir(const std::string &langDir);

    /**
     * @brief 加载资源
     * @param lang         : 语言
     * @param key          : 资源 key
     * @return 对应语言的 value 值，如果没找到返回空字符串
     */
    std::string Localize(const std::string & lang, const std::string & key) noexcept;

  private:
    std::unordered_map<std::string, std::shared_ptr<cpptoml::table>> res_;
    bool loaded_;
};

#endif // CODE_CPPEXAMPLE_03_I18N_I18N_RES_LOADER_H_