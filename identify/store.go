package identify

import (
    "github.com/dchest/captcha"
    "sync"
    "github.com/gin-gonic/gin"
    "errors"
    "time"
    "github.com/gin-contrib/sessions"
    "github.com/dazhenghu/ginApp/consts"
)

type sessionStore struct {
    captcha.Store
    contextIdMap map[string]*contextWithTime // 保存id与gin.context的映射关系，用于将id保存至session中
    expirePeriod time.Duration // 有效时长，单位秒，超过这个时长后过期删除

    sessionMutex sync.RWMutex
    memoryMutex sync.RWMutex
}

type contextWithTime struct {
    context *gin.Context
    createTime time.Time
}

var STORE_ERR_REMOVE_EMPTY = errors.New("invalid id, please reload")
var STORE_ERR_ID_EXISTS = errors.New("id exists, please reload")

var sessionStoreInstance *sessionStore
var once sync.Once

func GetSessionStore() *sessionStore {
    once.Do(func() {
        sessionStoreInstance = &sessionStore{
            contextIdMap: make(map[string]*contextWithTime),
        }
    })

    return sessionStoreInstance
}

func (ss *sessionStore) Init(expirePeriod time.Duration)  {
    sessStore := GetSessionStore()
    sessStore.expirePeriod = expirePeriod
}

/**
设置校验码，存储至session中
 */
func (ss *sessionStore) Set(id string, digits []byte) {
    ss.sessionMutex.Lock()
    defer ss.sessionMutex.Unlock()

    contextWithTime := ss.contextIdMap[id]
    sess := sessions.Default(contextWithTime.context)
    sess.Set(ss.keyByid(id), digits)
    sess.Save()
}

/**
从session中读取校验码
 */
func (ss *sessionStore) Get(id string, clear bool) (digits []byte)  {
    contextWithTime := ss.contextIdMap[id]
    sess := sessions.Default(contextWithTime.context) // 获取用户session

    overdue := ss.expirePeriod > 0 && contextWithTime.createTime.Add(ss.expirePeriod).Before(time.Now())
    if overdue || clear {
        // 到期了或者是需要删除的
        ss.sessionMutex.Lock()
        defer ss.sessionMutex.Unlock()

        if overdue {
            // 是过期
            digits = nil
        } else if clear {
            // 未过期但是要读取后删除，先读取
            digits = sess.Get(ss.keyByid(id)).([]byte)
        }
        // 删除存储在session中的对应id信息
        sess.Delete(ss.keyByid(id))
        sess.Save()
        // 删除内存中保存的context与id对应关系
        ss.RemoveContextId(contextWithTime.context, id)
        return
    }

    digits = sess.Get(ss.keyByid(id)).([]byte)

    return
}

/**
根据id生成对应key
 */
func (ss *sessionStore) keyByid(id string) string {
    return consts.SESSION_KEY_INDENTIFY + "_" + id
}

/**
在内存中添加context与id的映射数据
 */
func (ss *sessionStore) PushContextId(context *gin.Context, id string) error {
    ss.memoryMutex.Lock()
    defer ss.memoryMutex.Unlock()
    _, ok := ss.contextIdMap[id]
    if ok {
        return STORE_ERR_ID_EXISTS
    }
    ss.contextIdMap[id] = &contextWithTime{
        context: context,
        createTime: time.Now(),
    }
    return nil
}

/**
删除停留在内存中的context与id的映射数据
 */
func (ss *sessionStore) RemoveContextId(context *gin.Context, id string) error {
    ss.memoryMutex.Lock()
    defer ss.memoryMutex.Unlock()
    contWithTime, ok := ss.contextIdMap[id]
    if contWithTime.context != context {
        // 删除的数据对应不上
        return STORE_ERR_REMOVE_EMPTY
    }

    if !ok {
        // 本来就没有这个数据，则返回删除成功
        return nil
    }

    delete(ss.contextIdMap, id)
    return nil
}

