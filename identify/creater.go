package identify

import (
    "github.com/gin-gonic/gin"
    "github.com/dazhenghu/ginApp/safe/token"
    "github.com/dchest/captcha"
)

func New(context *gin.Context) string {
    return NewLen(context, captcha.DefaultLen)
}

func NewLen(context *gin.Context, length int) (id string) {
    id = randomId()
    sessStore := GetSessionStore()
    sessStore.PushContextId(context, id)
    sessStore.Set(id, captcha.RandomDigits(length))
    sessStore.RemoveContextId(context, id)
    return
}

func randomId() string {
    tokenObj := token.NewToken("")
    return tokenObj.GenerateToken()
}
