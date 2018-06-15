package session

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
    token2 "github.com/dazhenghu/ginApp/safe/token"
    "errors"
    "github.com/dazhenghu/ginApp/types"
)


func GenerateSessionToken(c *gin.Context, key string) (token string, err error) {
    tokenObj := token2.NewToken("")
    token = tokenObj.GenerateToken()
    session := sessions.Default(c)
    tokens := session.Get(key)
    var tokenList types.SliceString
    if tokens != nil {
        tokenList = types.NewSliceStringFromSlice(tokens.([]string))
    } else {
        tokenList = types.NewSliceString()
    }

    (&tokenList).Append(token)
    session.Set(key, tokenList.ToSlice())
    err = session.Save()
    return
}

func CheckSessionToken(c *gin.Context, key string, token string) (err error) {
    session := sessions.Default(c)
    tokens := session.Get(key)
    var tokenList types.SliceString
    if tokens != nil {
        tokenList = types.NewSliceStringFromSlice(tokens.([]string))
    } else {
        tokenList = types.NewSliceString()
    }

    removedIdx := (&tokenList).Remove(token)

    if removedIdx < 0 {
        err = errors.New("invalid token")
    } else {
        session.Set(key, tokenList.ToSlice())
        err = session.Save()
    }
    return
}
