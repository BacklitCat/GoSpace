package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"math/rand"
	"time"
	"unsafe"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// RandStringBytesMaskImprSrcUnsafe 高效生成随机字符串
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
// https://cloud.tencent.com/developer/article/2022005
func RandStringBytesMaskImprSrcUnsafe(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// HashSalt 用于对字符串进行加盐哈希
// str 是要加密的字符串
// salt 是随机生成的盐值
// 返回值是加盐哈希后的十六进制字符串
func HashSalt(str, salt string) (string, error) {
	// 创建一个 md5 哈希对象
	m := md5.New()
	// 将 str 写入哈希对象
	_, err := io.WriteString(m, str)
	if err != nil {
		return "", err
	}
	// 将 salt 写入哈希对象
	_, err = io.WriteString(m, salt)
	if err != nil {
		return "", err
	}
	// 计算哈希值，并转换成十六进制字符串
	return hex.EncodeToString(m.Sum(nil)), nil
}
