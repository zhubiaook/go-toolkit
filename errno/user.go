package errno

var (
	// ErrUserNotFound 表示未找到用户.
	ErrUserNotFound = &Errno{HTTP: 404, ErrorCode: 10103, ResponseCode: "ResourceNotFound.UserNotFound", Message: "User was not found."}

	// ErrPasswordIncorrect 表示密码不正确.
	ErrPasswordIncorrect = &Errno{HTTP: 401, ErrorCode: 10302, ResponseCode: "InvalidParameter.PasswordIncorrect", Message: "Password was incorrect."}

	// ErrPasswordIncorrect 表示用户名或密码不正确.
	ErrUserOrPasswordIncorrect = &Errno{HTTP: 401, ErrorCode: 10303, ResponseCode: "InvalidParameter.UserOrPasswordIncorrect", Message: "User or password was incorrect."}
)
