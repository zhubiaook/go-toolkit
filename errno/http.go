package errno

var (
	// OK 代表请求成功.
	OK = &Errno{HTTP: 200, ErrorCode: 0, ResponseCode: "OK", Message: "Response OK"}

	// InternalServerError 表示所有未知的服务器端错误.
	InternalServerError = &Errno{HTTP: 500, ErrorCode: 10000, ResponseCode: "InternalError", Message: "Internal server error."}

	// ErrPageNotFound 表示路由不匹配错误.
	ErrPageNotFound = &Errno{HTTP: 404, ErrorCode: 10101, ResponseCode: "ResourceNotFound.PageNotFound", Message: "Page not found."}

	// ErrMethodNotAllowed 表示Method不匹配错误
	ErrMethodNotAllowed = &Errno{HTTP: 405, ErrorCode: 10102, ResponseCode: "ResourceNotFound.MethodNotAllowed", Message: "Method not allowed."}

	// ErrInvalidParameter 表示所有验证失败的错误.
	ErrInvalidParameter = &Errno{HTTP: 400, ErrorCode: 10300, ResponseCode: "InvalidParameter", Message: "Parameter verification failed."}

	// ErrBind 表示参数绑定错误.
	ErrBind = &Errno{HTTP: 400, ErrorCode: 10301, ResponseCode: "InvalidParameter.BindError", Message: "Error occurred while binding the request body to the struct."}

	// ErrSignToken 表示签发 JWT Token 时出错.
	ErrSignToken = &Errno{HTTP: 401, ErrorCode: 10401, ResponseCode: "AuthFailure.SignTokenError", Message: "Error occurred while signing the JSON web token."}

	// ErrTokenInvalid 表示 JWT Token 格式错误.
	ErrTokenInvalid = &Errno{HTTP: 401, ErrorCode: 10402, ResponseCode: "AuthFailure.TokenInvalid", Message: "Token was invalid."}

	// ErrUnauthorized 表示请求没有被授权.
	ErrUnauthorized = &Errno{HTTP: 401, ErrorCode: 10403, ResponseCode: "AuthFailure.Unauthorized", Message: "Unauthorized."}
)
