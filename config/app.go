package config

/**
app配置
 */
type AppConfig struct {
    Addr        string             `yaml:"addr"`          // 监听端口
    Pulic       StaticConfig       `yaml:"public"`        // 静态资源路径
    ViewBaseDir string             `yaml:"viewBaseDir"`   // view路径
    Dblist      map[string]DbConfg `yaml:"dblist"`        // 数据库连接列表
}

/**
静态资源配置
 */
type StaticConfig struct {
    RelativePath string `yaml:"relativePath"`
    Root         string `yaml:"root"`
}
