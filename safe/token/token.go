package token

import (
    "time"
    "math/rand"
    "crypto/md5"
    "encoding/hex"
)

type TokenInterface interface {
    GenerateToken() string // 生成token
}

type Token struct {
    Prefix string // 前缀
}

/**
token对象
 */
func NewToken(prefix string) *Token {
    return &Token{
        Prefix: prefix,
    }
}

/**
生成token
 */
func (t *Token) GenerateToken() string {
    now := time.Now()
    str := t.Prefix + string(now.UnixNano()) + string(rand.Int()) + string(rand.Int())
    h := md5.New()
    h.Write([]byte(str))
    return hex.EncodeToString(h.Sum(nil))
}
