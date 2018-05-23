package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/dazhenghu/ginApp"
    "reflect"
    "fmt"
    "strings"
    "runtime"
)

type ControllerInterface interface {
    beforeAction(context *gin.Context) error
    afterAction(context *gin.Context) error
}


type Controller struct {
    ctrlName string
    child ControllerInterface
}

/**
controller初始化
 */
func (c *Controller) Init() error {
    // 通过反射注册action方法
    //cReflect := reflect.ValueOf(c)
    // 获取方法数量
    //numMethod := cReflect.NumMethod()

    // 读取当前类名，去掉controller后缀
    //fmt.Printf("%+v\n%+v\n%+v\n%+v\n", pc,file,line,ok)
    ctrlType :=  reflect.TypeOf(c).String()
    fmt.Printf("type:%+v\n", ctrlType)
    ctrlType = ctrlType[strings.LastIndex(ctrlType, "."):]
    c.ctrlName = strings.TrimRight(ctrlType, "Controller")

    // 获取上层caller，即子类的文件名
    pc, file, _, ok := runtime.Caller(1)
    if ok {
        c.ctrlName = file[strings.LastIndex(file, "/") + 1:strings.LastIndex(file, ".")]
        fmt.Printf("ctrlName:%s\n", c.ctrlName)
        fmt.Printf("func:%+v\n", runtime.FuncForPC(pc).Name())
    }

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
        // 启动协程执行action
        go func(c *Controller, context *gin.Context, handler gin.HandlerFunc) {
            // 执行handler之前执行beforeAction
            berforeErr := c.child.beforeAction(context)
            if berforeErr != nil {
                panic(berforeErr)
            }
            handler(context)
            // 执行handler之后执行beforeAction
            afterErr := c.child.afterAction(context)
            if afterErr != nil {
                panic(afterErr)
            }
        }(c, context, handler)

    }
}