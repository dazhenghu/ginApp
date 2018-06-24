package session

import (
    "github.com/gin-contrib/sessions"
    "github.com/dazhenghu/ginApp/config"
    "strconv"
    "errors"
    "fmt"
)

const (
    SESSION_TYPE_COOKIE    string = "cookie"    //
    SESSION_TYPE_MEMCACHED string = "memcached" // 保存在memecache中
    SESSION_TYPE_REDIS     string = "redis"     // 保存在redis中
    SESSION_TYPE_MONGO     string = "mongo"     // 保存在mongo中
)

var DefaultKey string = "ginapp-session" // cookie中保存session的key

func NewStore(appConfig *config.AppConfig) (store sessions.Store, err error) {

    secret := appConfig.Secret
    if secret == "" {
        secret = "ZJuZZHwk626kwcHI2"
    }
    sessionConf := appConfig.SessionCnf.ConnectCnf // 配置中读取的session配置信息

    options := sessions.Options{
        Path: "/",
        HttpOnly: true,
    }

    switch appConfig.SessionCnf.Type {
    case SESSION_TYPE_COOKIE:
        store = sessions.NewCookieStore([]byte(secret))
        store.Options(options)
        return
    case SESSION_TYPE_REDIS:
        size, _ := strconv.Atoi(sessionConf["size"])
        network := sessionConf["network"]
        adress := sessionConf["adress"]
        password := sessionConf["password"]
        store, err = sessions.NewRedisStore(size, network, adress, password, []byte(secret))
        store.Options(options)
        return
    }

    err = errors.New(fmt.Sprintf("not valid session type:%s", appConfig.SessionCnf.Type))

    return
}
