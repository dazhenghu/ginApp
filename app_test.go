package ginApp

import (
    "testing"
    "github.com/dazhenghu/util/fileutil"
    "path"
    "fmt"
    "time"
    "strconv"
    "github.com/gin-gonic/gin"
    "github.com/dazhenghu/ginApp/identify"
)

var App *GinApp

func init()  {
    ginAppInstance := Instance()
    // 保存全局app实例
    App = ginAppInstance
    // 获取挡墙文件夹
    currPath, _ := fileutil.GetCurrentDirectory()
    // 设置common文件夹位置
    App.SetCommonPath(path.Join(currPath, "./common"))
    // 设置根目录文件夹位置
    App.SetRootPath(currPath)
    // 读取默认位置的配置文件
    App.DefaultLoadConfig("")
    fmt.Printf("app config:%+v\n", App.AppConfig)
    // 初始化session配置
    sessionErr := App.InitSession()
    fmt.Printf("session err:%+v\n", sessionErr)

    // 初始化验证码模块，过期时间为10分钟
    App.InitIdentify(10 * time.Minute)
}


func TestGinApp_InitIdentify(t *testing.T) {

    //digits := []byte("123") //[]byte{1,2,3}
    digits := "789"
    for index, n := range digits {
        val, _ := strconv.Atoi(string(n))
        fmt.Printf("index:%d n:%d\n",  index, val)
    }

    App.Engine().GET("login", func(context *gin.Context) {
        id := identify.New(context)
        fmt.Printf("id:%s\n", id)
        return
    })

    App.Engine().GET("captcha/:name", func(context *gin.Context) {
        name := context.Param("name")
        fmt.Printf("name:%+v\n", name)
        //captchaId := captcha2.New()
        //fmt.Printf("captchaId:%+v\n", captchaId)
        captcha := identify.CaptchaNew(240, 80)
        err := captcha.Handle(context)
        fmt.Printf("err:%+v\n", err)
        return
    })

    App.Run(":8888")
}
