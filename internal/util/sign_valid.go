package util

import (
	"net"
	"regexp"
	"strings"
)

var (
	// 匹配邮箱的正则表达式
	emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	// 匹配中国大陆手机号的正则表达式
	phoneRegex = regexp.MustCompile(`^1[3-9]\d{9}$`)
)

// IsEmailValid 检查邮件地址是否合法
func IsEmailValid(e *string) bool {
	if len(*e) < 3 && len(*e) > 254 {
		return false
	}
	if !emailRegex.MatchString(*e) {
		return false
	}
	parts := strings.Split(*e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}
	return true
}

// IsPhoneValid 检查手机号是否合法
func IsPhoneValid(p *string) bool {
	return phoneRegex.MatchString(*p)
}
