package errno

// SQL错误
const (
	ErrSQL = 1001000 + iota
	ErrSQLSyntax
	ErrSQLQueryEmpty
)

// 注册用户阶段错误
const (
	ErrSignUp = 2000000 + iota
	ErrSignUpUnknownEmail
	ErrSignUpUnknownPhone
	ErrSignUpUnknownEmailOrPhone
	ErrSignUpEmailOrPhoneDuplicate
)

// 查询用户信息错误
const (
	ErrSelectUser = 2001000 + iota
	ErrSelectUserByEmailEmpty
	ErrSelectUserByPhoneEmpty
	ErrSelectUserUnknownEmailOrPhone
)

var ErrMsgMap = map[int]string{
	ErrSQL:           "SQL错误",
	ErrSQLSyntax:     "SQL拼写错误",
	ErrSQLQueryEmpty: "SQL查询为空",

	ErrSignUp:                      "注册用户阶段错误",
	ErrSignUpUnknownEmail:          "注册账号需要邮箱",
	ErrSignUpUnknownPhone:          "注册账号需要手机号",
	ErrSignUpUnknownEmailOrPhone:   "注册账号需要邮箱或手机号",
	ErrSignUpEmailOrPhoneDuplicate: "您使用的手机号或邮箱已被注册",

	ErrSelectUser:                    "查询用户信息错误",
	ErrSelectUserByEmailEmpty:        "通过该邮箱查询用户结果为空",
	ErrSelectUserByPhoneEmpty:        "通过该手机查询用户结果为空",
	ErrSelectUserUnknownEmailOrPhone: "查询用户信息需要邮箱或手机号",
}
