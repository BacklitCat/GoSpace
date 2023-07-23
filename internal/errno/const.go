package errno

// util错误
const (
	ErrUtil = 1001000 + iota
	ErrUtilHashSalt
)

// SQL错误
const (
	ErrSQL = 1002000 + iota
	ErrSQLSyntax
	ErrSQLQueryEmpty
)

// 用户Email/Phone错误
const (
	ErrEmailOrPhone = 2001000 + iota
	ErrUnknownEmail
	ErrUnknownPhone
	ErrUnknownEmailAndPhone
	ErrDuplicateEmailOrPhone
	ErrNotExistEmail
	ErrNotExistPhone
	ErrNotExistEmailAndPhone
)

// 注册用户阶段错误
const (
	ErrSignUp = 2002000 + iota
	ErrSignUpTooMany
)

// 查询用户信息错误
const (
	ErrSelectUser = 2003000 + iota
	ErrSelectUserByIdEmpty
	ErrSelectUserByEmailEmpty
	ErrSelectUserByPhoneEmpty
)

// 用户登陆阶段错误
const (
	ErrSignIn = 2004000 + iota
	ErrSignInWrongPwd
	ErrSignInMaxRetries
	ErrSignInUnsafe
	ErrSignInDisabledStatus
)

var ErrMsgMap = map[int]string{
	ErrUtil:         "内部util包错误",
	ErrUtilHashSalt: "计算哈希密码时出现错误",

	ErrSQL:           "SQL错误",
	ErrSQLSyntax:     "SQL拼写错误",
	ErrSQLQueryEmpty: "SQL查询为空",

	ErrEmailOrPhone:          "用户Email/Phone错误",
	ErrUnknownEmail:          "Email参数为空",
	ErrUnknownPhone:          "Phone参数为空",
	ErrUnknownEmailAndPhone:  "需要邮箱或手机号",
	ErrDuplicateEmailOrPhone: "邮箱或手机号已经注册",
	ErrNotExistEmail:         "该邮箱不存在",
	ErrNotExistPhone:         "该手机号不存在",
	ErrNotExistEmailAndPhone: "邮箱或手机号不存在",

	ErrSignUp:        "注册用户阶段错误",
	ErrSignUpTooMany: "注册次数过多",

	ErrSelectUser:             "查询用户信息错误",
	ErrSelectUserByIdEmpty:    "通过该ID查询用户结果为空",
	ErrSelectUserByEmailEmpty: "通过该邮箱查询用户结果为空",
	ErrSelectUserByPhoneEmpty: "通过该手机查询用户结果为空",

	ErrSignIn:               "用户登陆阶段错误",
	ErrSignInWrongPwd:       "密码错误",
	ErrSignInMaxRetries:     "已达到最大登陆次数",
	ErrSignInUnsafe:         "该次登陆有风险",
	ErrSignInDisabledStatus: "账户被封禁",
}
