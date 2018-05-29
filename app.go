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
    rootPath string // 应用根目录
    engine *gin.Engine
    envMode string // 当前环境
    AppConfig *config.AppConfig // app的配置信息
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
                AppConfig:&config.AppConfig{},
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

/**
设置项目根目录，通常是web所在目录
 */
func (app *GinApp)SetRootPath(root string)  {
    app.rootPath = root
}

/**
读取项目根目录
 */
func (app *GinApp)GetRootPath() string {
    return app.rootPath
}

func (app *GinApp)DefaultLoadConfig(configDirPath string)  {
    if configDirPath == "" {
        // 未设置路径
        configDirPath = app.GetRootPath()
    }

    config.DefaultLoadFromYaml(configDirPath, app.AppConfig)
}