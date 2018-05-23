package ginApp

import (
    "github.com/gin-gonic/gin"
    "sync"
)

type GinApp struct {
    engine *gin.Engine
    appConfig map[string]interface{} // app的配置信息
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
                appConfig:make(map[string]interface{}),
            }
        }
    }

    return instance
}

func Run(addr ...string) *GinApp {
    Instance().engine.Run(addr...)
    return Instance()
}

func (app *GinApp)Engine() *gin.Engine  {
    return app.engine
}