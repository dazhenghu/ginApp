package config

import (
    "runtime"
    "path/filepath"
    "errors"
    "github.com/dazhenghu/ginApp/consts"
)

var EnvMode string = consts.ENV_DEBUG
var AppLibAbsPath string

func init()  {
    AppLibAbsPath, _= getAppLibAbsPath()
}

/**
app配置
 */
type AppConfig struct {
    Addr          string             `yaml:"addr"`        // 监听端口
    Pulic         StaticConfig       `yaml:"public"`      // 静态资源路径
    Secret        string             `yaml:"secret"`      // 用于系统加密的秘钥
    ViewBaseDir   string             `yaml:"viewBaseDir"` // view路径
    Dblist        map[string]DbConfg `yaml:"dblist"`      // 数据库连接列表
    AppLibAbsPath string                                  // ginapp所在的绝对路径
    SessionCnf    SessionConfig      `yaml:"session"`     // session配置
}

/**
静态资源配置
 */
type StaticConfig struct {
    RelativePath string `yaml:"relativePath"`
    Root         string `yaml:"root"`
}

/**
读取
 */
func getAppLibAbsPath() (path string, err error) {
    _, file, _, ok := runtime.Caller(0)
    if !ok {
        err = errors.New("runtime get caller err")
        return
    }

    // ginapp所在的路径应该是config的上层目录，所以此处要加【../】
    path, err = filepath.Abs(filepath.Join(filepath.Dir(file), "../"))
    if err != nil {
        return
    }

    return
}
