#ifndef CODE_CPPEXAMPLE_03_I18N_I18N_ERROR_H_
#define CODE_CPPEXAMPLE_03_I18N_I18N_ERROR_H_

#include <exception>
#include <stdint.h>
#include <string>

/**
 * @brief 符合 RESTful 规范的异常基类
 */
class HttpRestException : public std::exception {
  public:
    HttpRestException(std::string module, std::string resouceKey) : module(module), resouceKey(resouceKey) {}
    HttpRestException(std::string module, std::string resouceKey, int32_t errorCode)
        : module(module), resouceKey(resouceKey), errorCode(errorCode) {}

    HttpRestException(const std::string &codeFileName, int32_t codeLine, std::string module, std::string resouceKey)
        : codeFileName(codeFileName), codeLine(codeLine), module(module), resouceKey(resouceKey) {}
    HttpRestException(const std::string &codeFileName, int32_t codeLine, std::string module, std::string resouceKey,
                      int32_t errorCode)
        : codeFileName(codeFileName), codeLine(codeLine), module(module), resouceKey(resouceKey), errorCode(errorCode) {
    }
    ~HttpRestException() = default;

  public:
    std::string GetModule() const { return module; }
    std::string GetResouceKey() const { return resouceKey; }
    int32_t ErrorCode() const { return errorCode; }
    std::string CodeFileName() const { return codeFileName; }
    int32_t CodeLine() const { return codeLine; };

  private:
    std::string module;
    std::string resouceKey;
    int32_t errorCode;
    std::string codeFileName;
    int32_t codeLine;
};

/**
 * @brief 客户端类错误
 */
class ClientException : public HttpRestException {
  public:
    using HttpRestException::HttpRestException;
};

/**
 * @brief 服务端类错误
 */
class ServerException : public HttpRestException {
  public:
    using HttpRestException::HttpRestException;
};

//
//
// 4xx 错误
//
//

/**
 * @brief 400 Bad Request
 * 由于明显的客户端错误(如：参数格式错误，未提供所需参数，参数为空，参数超出范围等)而导致API请求失败，返回由此状态码构造的错误码
 */
class BadRequestException : public ClientException {
  public:
    using ClientException::ClientException;
};

/**
 * @brief 401 Unauthorized
 * 由于身份认证失败(如：token过期或不存在等)而导致API请求失败，返回由此状态码构造的错误码
 */
class UnauthorizedException : public ClientException {
  public:
    using ClientException::ClientException;
};

/**
 * @brief 403 Forbidden
 * 由于逻辑错误、权限错误、不正确的操作等原因而导致API请求失败，返回由此状态码构造的错误码
 */
class ForbiddenException : public ClientException {
  public:
    using ClientException::ClientException;
};

/**
 * @brief 404 Not Found
 * 由于服务器未找到 Request-URI 所指向的数据而导致API请求失败，返回由此状态码构造的错误码
 */
class NotFoundException : public ClientException {
  public:
    using ClientException::ClientException;
};

/**
 * @brief 409 Conflict
 * 由于与目标资源的当前状态相冲突，从而导致API请求失败，返回由此状态码构造的错误码
 */
class ConflictException : public ClientException {
  public:
    using ClientException::ClientException;
};

//
//
// 5xx 错误
//
//

/**
 * @brief 500 Internal Server Error
 * 由于服务器内部环境或逻辑错误而导致API请求失败，返回由此状态码构造的错误码
 */
class InternalServerErrorException : public ServerException {
  public:
    using ServerException::ServerException;
};

/**
 * @brief 501 Not Implemented
 * 服务器无法识别请求方法并且不支持任何资源，从而导致API请求失败，返回由此状态码构造的错误码
 */
class NotImplementedException : public ServerException {
  public:
    using ServerException::ServerException;
};

/**
 * @brief 503 Service Unavailable
 * 由于服务器暂时不可用而导致API请求失败，返回由此状态码构造的错误码
 */
class ServiceUnavailableException : public ServerException {
  public:
    using ServerException::ServerException;
};

#endif // CODE_CPPEXAMPLE_03_I18N_I18N_ERROR_H_