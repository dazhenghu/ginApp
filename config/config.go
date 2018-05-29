package config

import (
    "github.com/dazhenghu/ginApp"
    "github.com/dazhenghu/util/fileutil"
    "path"
    "github.com/go-yaml/yaml"
    "io/ioutil"
    "fmt"
)

func DefaultLoadFromYaml(configDirPath string) {
    app := ginApp.Instance()
    if configDirPath == "" {
        // 未设置路径
        configDirPath = app.GetRootPath()
    }

    mainConfigPath := path.Join(configDirPath, "main.yaml")
    if exists, err := fileutil.PathExists(mainConfigPath); exists && err == nil {
        // main配置存在
        // 读取配置文件
        configFile, err := ioutil.ReadFile(mainConfigPath)
        if err != nil {
            panic(fmt.Sprintf("load config err:%+v\n", err))
        }
        err = yaml.Unmarshal(configFile, app.AppConfig)
        if err != nil {
            panic(fmt.Sprintf("load config unmarshal err:%+v\n", err))
        }
    }
}


func LoadFromYaml(filePath string)  {
    
}
