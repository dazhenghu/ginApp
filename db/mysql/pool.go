package mysql

import (
    "github.com/dazhenghu/ginApp/db"
    "github.com/jinzhu/gorm"
    "sync"
    "errors"
)

type MysqlConnectPool struct {
    db.DbConnectPool
    poolSlice []*gorm.DB // 存储链接的切片
    lastPopIndex int // 最后一次弹出链接索引
    sync.RWMutex
}

/**
生成连接池
 */
func NewConnectPool(len int, dbConfig map[string]string) *MysqlConnectPool  {
    mysqlConnectPool := &MysqlConnectPool{}
    mysqlConnectPool.Init(len, dbConfig)
    return mysqlConnectPool
}

func (pool *MysqlConnectPool)Init(len int, dbConfig map[string]string) error  {
    pool.poolSlice = make([]*gorm.DB, len, len * 2)

    dsn, ok := dbConfig["dsn"]
    if !ok {
        return errors.New("mysql connect pool init error, dsn is empty")
    }

    for i := 0; i < len; i++  {
        dbInstance, err := gorm.Open("mysql", dsn)
        if err != nil {
            return err
        }
        pool.Push(dbInstance)
    }

    return nil
}

func (pool *MysqlConnectPool)Push(db *gorm.DB) error  {
    pool.Lock()
    defer pool.Unlock()
    pool.poolSlice = append(pool.poolSlice, db)
    return nil
}

func (pool *MysqlConnectPool)Pop() (*gorm.DB, error)  {
    pool.RLock()
    defer pool.RUnlock()

    // 轮询方式获取链接
    if pool.lastPopIndex + 1 <= len(pool.poolSlice) - 1 {
        pool.lastPopIndex++
    } else {
        pool.lastPopIndex = 0
    }

    return pool.poolSlice[pool.lastPopIndex], nil
}



