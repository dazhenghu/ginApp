package session

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
    token2 "github.com/dazhenghu/ginApp/safe/token"
    "errors"
    "container/list"
)

func GenerateSessionToken(c *gin.Context, key string) (token string, err error) {
    tokenObj := token2.NewToken("")
    token = tokenObj.GenerateToken()
    session := sessions.Default(c)
    tokenList := list.New()
    tokenList.PushBack(token)
    tokenList.PushBack("asdasd")
    session.Set(key, tokenList)
    err = session.Save()
    return
}

func CheckSessionToken(c *gin.Context, key string, token string) (err error) {
    session := sessions.Default(c)
    sessionTokens := session.Get(key).([]string)
    for _, val := range sessionTokens {
        if val == token {

            return
        }
    }

    err = errors.New("invalid token")
    return
}
