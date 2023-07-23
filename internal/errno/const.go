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

// 注册用户阶段错误
const (
	ErrSignUp = 2000000 + iota
	ErrSignUpUnknownEmail
	ErrSignUpUnknownPhone
	ErrSignUpUnknownEmailAndPhone
	ErrSignUpEmailOrPhoneDuplicate
)

// 查询用户信息错误
const (
	ErrSelectUser = 2001000 + iota
	ErrSelectUserById
	ErrSelectUserByEmailEmpty
	ErrSelectUserByPhoneEmpty
	ErrSelectUserUnknownEmailAndPhone
)

// 用户登陆阶段错误
const (
	ErrSignIn = 2002000 + iota
	ErrSignInUnknownEmailAndPhone
	ErrSignInNotExistEmailAndPhone
	ErrSignInNotExistEmail
	ErrSignInNotExistPhone
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

	ErrSignUp:                      "注册用户阶段错误",
	ErrSignUpUnknownEmail:          "注册账号需要邮箱",
	ErrSignUpUnknownPhone:          "注册账号需要手机号",
	ErrSignUpUnknownEmailAndPhone:  "注册账号需要邮箱或手机号",
	ErrSignUpEmailOrPhoneDuplicate: "您使用的手机号或邮箱已被注册",

	ErrSelectUser:                     "查询用户信息错误",
	ErrSelectUserById:                 "通过该ID查询用户结果为空",
	ErrSelectUserByEmailEmpty:         "通过该邮箱查询用户结果为空",
	ErrSelectUserByPhoneEmpty:         "通过该手机查询用户结果为空",
	ErrSelectUserUnknownEmailAndPhone: "查询用户信息需要邮箱或手机号",

	ErrSignIn:                      "用户登陆阶段错误",
	ErrSignInUnknownEmailAndPhone:  "登陆账号需要邮箱或手机号",
	ErrSignInNotExistEmailAndPhone: "邮箱或手机号不存在",
	ErrSignInNotExistEmail:         "该邮箱不存在",
	ErrSignInNotExistPhone:         "该手机号不存在",
	ErrSignInWrongPwd:              "密码错误",
	ErrSignInMaxRetries:            "已达到最大登陆次数",
	ErrSignInUnsafe:                "该次登陆有风险",
	ErrSignInDisabledStatus:        "账户被封禁",
}
