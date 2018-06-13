package token

import (
    "time"
    "math/rand"
    "crypto/md5"
    "encoding/hex"
    "strconv"
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
    randObj := rand.New(rand.NewSource(now.UnixNano()))
    str := t.Prefix + strconv.FormatInt(now.UnixNano(), 10) + strconv.Itoa(randObj.Int()) + strconv.Itoa(randObj.Int())
    h := md5.New()
    h.Write([]byte(str))
    return hex.EncodeToString(h.Sum(nil))
}
