package config

type SessionConfig struct {
    Type       string            `yaml:"type"`   // 数据库类型，如：redis
    ConnectCnf map[string]string `yaml:"config"` // 配置信息
}
