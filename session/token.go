package session

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
    token2 "github.com/dazhenghu/ginApp/safe/token"
)

func GenerateSessionToken(c *gin.Context, key string) (token string) {
    tokenObj := token2.NewToken("")
    token = tokenObj.GenerateToken()
    session := sessions.Default(c)
    session.Set(key, token)
    return
}
