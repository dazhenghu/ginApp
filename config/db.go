package config

type DbConfg struct {
    Type string `yaml:"type"` // 数据库类型，如：mysql
    Dsn  string `yaml:"dsn"`  // 数据库连接
}

