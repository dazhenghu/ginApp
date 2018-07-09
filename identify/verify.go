package identify

import (
    "github.com/dchest/captcha"
    "github.com/gin-gonic/gin"
    "github.com/dazhenghu/ginApp/logs"
    "fmt"
)

func Verify(id string, digits []byte) bool  {
    return captcha.Verify(id, digits)
}

func VerifyString(id string, digits string, context *gin.Context) bool {
    sessStore := GetSessionStore()
    err := sessStore.PushContextId(context, id) // 注入context与id的关系，主要用于SessionStore将id存入到session中
    if err != nil {
        logs.Error(fmt.Sprintf("captcha verify err:%+v\n", err))
        return false;
    }
    res := captcha.VerifyString(id, digits)
    sessStore.RemoveContextId(id)
    return res
}
