package db

import "github.com/jinzhu/gorm"

/**
数据库连接池
 */
type DbConnectPool interface {
    Init(len int, dbConfig map[string]string) error // 连接池初始化，len：链接缓存个数
    Push(*gorm.DB) error // 往连接池中添加连接
    Pop() (*gorm.DB, error) // 获取链接
}

