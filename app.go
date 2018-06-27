package ginApp

import (
    "github.com/gin-gonic/gin"
    "sync"
    "github.com/dazhenghu/ginApp/config"
    "path"
    "github.com/dazhenghu/util/fileutil"
    "github.com/jinzhu/gorm"
    "fmt"
    "html/template"
    "github.com/dazhenghu/util/htmlutil"
    "github.com/gin-contrib/sessions"
    "github.com/dazhenghu/ginApp/session"
    "time"
    "github.com/dazhenghu/ginApp/identify"
    "github.com/dchest/captcha"
)

const (
    // 当前环境
    ENV_DEBUG string = "debug"
    ENV_TEST string = "test"
    ENV_PROD string = "prod"
)

type GinApp struct {
    rootPath   string // 应用根目录
    commonPath string // common目录
    engine     *gin.Engine
    envMode    string            // 当前环境
    AppConfig  *config.AppConfig // app的配置信息
    Db *gorm.DB // 默认数据库，主库
}

var instance *GinApp
var mutex sync.Mutex
var dbOnce sync.Once

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

    // 初始化一些常用方法到模板中
    tmplateFuncMap := initFuncMap()
    instance.Engine().SetFuncMap(tmplateFuncMap)

    // 初始化session配置


    return instance
}

func initFuncMap() template.FuncMap {
    tmplateFuncMap := template.FuncMap{
        "unescape": htmlutil.Unescape,
    }

    return tmplateFuncMap
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

/**
设置common文件夹
 */
func (app *GinApp)SetCommonPath(commonPath string)  {
    app.commonPath = commonPath
}

func (app *GinApp)GetCommonPath() string  {
    return app.commonPath
}

func (app *GinApp)DefaultLoadConfig(configDirPath string)  {
    if app.commonPath != "" {
        if exists, err := fileutil.PathExists(app.commonPath); exists && err == nil {
            commonConfigDirPath := path.Join(app.commonPath, "config")
            // 先读取common中的配置
            config.DefaultLoadFromYaml(commonConfigDirPath, app.AppConfig)
        }
    }

    if configDirPath == "" {
        // 未设置路径
        configDirPath = path.Join(app.GetRootPath(), "config")
    }

    config.DefaultLoadFromYaml(configDirPath, app.AppConfig)
}

/**
获取默认db
 */
func (app *GinApp) GetDb() *gorm.DB {
    dbOnce.Do(func() {
        dbConfig, ok := app.AppConfig.Dblist["db"]
        if !ok {
            return
        }

        db, err := gorm.Open(dbConfig.Type, dbConfig.Dsn)
        if err != nil {
            panic(fmt.Sprintf("default db init err,err:%+v", err))
        }
        app.Db = db
    })

    return app.Db
}

/**
初始化session配置
 */
func (app *GinApp) InitSession() error {
    store, err := session.NewStore(app.AppConfig)
    app.Engine().Use(sessions.Sessions(session.DefaultKey, store))
    return err
}

/**
初始化验证码模块
 */
func (app *GinApp) InitIdentify(expirePeriod time.Duration)  {
    identifyStore := identify.GetSessionStore(expirePeriod)
    captcha.SetCustomStore(identifyStore)
}
