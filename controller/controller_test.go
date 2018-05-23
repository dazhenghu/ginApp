package controller

import (
    "testing"
    "github.com/gin-gonic/gin"
    "fmt"
    "time"
)

type TmpController struct {
    Controller
}

/**
action调用前回调
 */
func (c *TmpController) beforeAction(context *gin.Context) error {
    fmt.Printf("TmpController:%s\n", "before action")
    return nil
}

/**
action调用后回调
 */
//func (c *TmpController) afterAction(context *gin.Context) error  {
//    fmt.Printf("TmpController:%s\n", "after action")
//    return nil
//}

func TestController_Init(t *testing.T) {
    ctrl := &TmpController{}
    ctrl.Init(ctrl)
    f := ctrl.hook(func(context *gin.Context) {
        fmt.Printf("invok:%s\n", "invoke")
    })

    f(&gin.Context{})

    time.Sleep(2 * time.Second)
}
