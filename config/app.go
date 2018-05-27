package config

type AppConfig struct {
    Addr        string             `yaml:"addr"`          // 监听端口
    Pulic       string             `yaml:"public"`        // 静态资源路径
    ViewBaseDir string             `yaml:"viewBaseDir"`   // view路径
    Dblist      map[string]DbConfg `yaml:"dblist"`        // 数据库连接列表
}
