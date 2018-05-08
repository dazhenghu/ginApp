package controller

import (
    "github.com/gin-gonic/gin"
    appError "github.com/dazhenghu/ginApp/error"
)

type Controller struct {
    RouterGroup gin.RouterGroup
}

/**
action调用前回调
 */
func (c *Controller)beforeAction(context *gin.Context) appError.BeforeActionErr {
    return nil
}

/**
action调用后回调
 */
func (c *Controller)afterAction(context *gin.Context) appError.AfterActionErr  {
    return nil
}

/**
GET方法路由handle设置
 */
func (c *Controller) Get(relativePath string, handler gin.HandlerFunc) {
    c.RouterGroup.GET(relativePath, c.actionHook(handler))
}

/**
POST方法路由handle设置
 */
func (c *Controller) Post(relativePath string, handler gin.HandlerFunc)  {
    c.RouterGroup.POST(relativePath, c.actionHook(handler))
}

func (c *Controller) actionHook(handler gin.HandlerFunc) func(context *gin.Context)  {
    return func(context *gin.Context) {
        // 执行handler之前执行beforeAction
        berforeErr := c.beforeAction(context)
        if berforeErr != nil {
            panic(berforeErr)
        }
        handler(context)
        // 执行handler之后执行beforeAction
        afterErr := c.afterAction(context)
        if afterErr != nil {
            panic(afterErr)
        }
    }
}