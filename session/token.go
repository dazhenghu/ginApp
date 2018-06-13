package session

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
    token2 "github.com/dazhenghu/ginApp/safe/token"
    "errors"
)

func GenerateSessionToken(c *gin.Context, key string) (token string) {
    tokenObj := token2.NewToken("")
    token = tokenObj.GenerateToken()
    session := sessions.Default(c)
    session.Set(key, token)
    return
}

func CheckSessionToken(c *gin.Context, key string, token string) (err error) {
    session := sessions.Default(c)
    sessionToken := session.Get(key).(string)
    if sessionToken == "" {
        err = errors.New("token is empty")
    }
    if sessionToken != token {
        err = errors.New("invalid token")
    }
    return
}
