package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/dazhenghu/ginApp"
    "reflect"
    "fmt"
    "strings"
)

type ControllerInterface interface {
    beforeAction(context *gin.Context) error
    afterAction(context *gin.Context) error
}


type Controller struct {
    ctrlName string // controller 名
    viewDir string // view 默认文件夹（相对路径）
    this ControllerInterface // 实例
}

/**
controller初始化
 */
func (c *Controller) Init(this ControllerInterface) error {
    c.this = this
    // 读取当前类名，去掉controller后缀
    ctrlType :=  reflect.TypeOf(c.this).String()
    if ginApp.Instance().Mode() == ginApp.ENV_DEBUG {
        fmt.Printf("Controller Init, controller type:%+v\n", ctrlType)
    }
    ctrlType = ctrlType[strings.LastIndex(ctrlType, ".") + 1:]
    c.ctrlName = strings.TrimRight(ctrlType, "Controller")
    return nil
}

/**
action调用前回调
 */
func (c *Controller) beforeAction(context *gin.Context) error {
    fmt.Printf("Controller:%s\n", "before action")
    return nil
}

/**
action调用后回调
 */
func (c *Controller) afterAction(context *gin.Context) error  {
    fmt.Printf("Controller:%s\n", "after action")
    return nil
}

/**
GET方法路由handle设置
 */
func (c *Controller) Get(relativePath string, handler gin.HandlerFunc) {
    ginApp.Instance().Engine().GET(relativePath, c.hook(handler))
}

/**
POST方法路由handle设置
 */
func (c *Controller) Post(relativePath string, handler gin.HandlerFunc)  {
    ginApp.Instance().Engine().POST(relativePath, c.hook(handler))
}

func (c *Controller) hook(handler gin.HandlerFunc) func(context *gin.Context)  {
    return func(context *gin.Context) {
        berforeErr := c.this.beforeAction(context)
        if berforeErr != nil {
            panic(berforeErr)
        }
        handler(context)
        // 执行handler之后执行beforeAction
        afterErr := c.this.afterAction(context)
        if afterErr != nil {
            panic(afterErr)
        }
        // 启动协程执行action
        //go func(c *Controller, context *gin.Context, handler gin.HandlerFunc) {
        //    // 执行handler之前执行beforeAction
        //
        //}(c, context, handler)

    }
}