package ginApp

import (
    "github.com/gin-gonic/gin"
    "sync"
    "github.com/dazhenghu/ginApp/config"
)

const (
    // 当前环境
    ENV_DEBUG string = "debug"
    ENV_TEST string = "test"
    ENV_PROD string = "prod"
)

type GinApp struct {
    engine *gin.Engine
    envMode string // 当前环境
    appConfig *config.AppConfig // app的配置信息
}

var instance *GinApp
var mutex sync.Mutex

/**
app单例
 */
func Instance() *GinApp {
    if instance == nil {
        mutex.Lock()
        defer mutex.Unlock()
        if instance == nil {
            app := gin.Default()
            instance = &GinApp{
                engine:app,
                appConfig:&config.AppConfig{},
            }
        }
    }

    return instance
}

func (app *GinApp)Run(addr ...string) *GinApp {
    Instance().engine.Run(addr...)
    return Instance()
}

func (app *GinApp)Engine() *gin.Engine  {
    return app.engine
}

/**
设置当前运行环境：debug、test、release
 */
func (app *GinApp)SetMode(mode string)  {
    app.envMode = mode
}

func (app *GinApp)Mode() string  {
    return app.envMode
}