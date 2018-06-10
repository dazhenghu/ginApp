package config

import (
    "github.com/dazhenghu/util/fileutil"
    "path"
    "github.com/go-yaml/yaml"
    "io/ioutil"
    "fmt"
    "path/filepath"
)

func DefaultLoadFromYaml(configDirPath string, appConfig *AppConfig) {
    // 先加载库中的配置文件
    defaultConfig, _ := ioutil.ReadFile(filepath.Join(AppLibAbsPath, "common/config/main.yaml"))
    yaml.Unmarshal(defaultConfig, appConfig)

    mainConfigPath := path.Join(configDirPath, "main.yaml")
    if exists, err := fileutil.PathExists(mainConfigPath); exists && err == nil {
        // main配置存在
        // 读取配置文件
        configFile, err := ioutil.ReadFile(mainConfigPath)
        if err != nil {
            panic(fmt.Sprintf("load config err:%+v\n", err))
        }
        err = yaml.Unmarshal(configFile, appConfig)
        if err != nil {
            panic(fmt.Sprintf("load config unmarshal err:%+v\n", err))
        }
    }

    appConfig.AppLibAbsPath = AppLibAbsPath
}


func LoadFromYaml(filePath string)  {
    
}



