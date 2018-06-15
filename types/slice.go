package types

import "sync"

type SliceString struct {
    slice []string // 存储数据的切片
    maxLen int     // 允许的最大长度，超过在从左侧删除
    sync.Mutex
}

func NewSliceString() (ss *SliceString) {
    ss = &SliceString{
        slice: make([]string, 0),
        maxLen: 100, // 默认最大长度为100
    }
    return
}

func NewSliceStringFromSlice(arr []string) (ss *SliceString) {
    ss = &SliceString{
        slice: arr,
        maxLen:100, // 默认最大长度为100
    }
    return
}

func (ss *SliceString) Remove(rmVal string) (removedIdx int) {
    ss.Lock()
    defer ss.Unlock()
    removedIdx = -1
    for i, val := range ss.slice {
        if val == rmVal {
            removedIdx = i
            break
        }
    }

    if removedIdx > -1 {
        ss.slice = append(ss.slice[:removedIdx], ss.slice[removedIdx+1:]...)
    }
    return
}

func (ss *SliceString) Append(appendVal string)  {
    dataLen := len(ss.slice)
    if dataLen >= ss.maxLen {
        // 达到最大长度要求时去掉头部第一个元素再追加
        ss.slice = append(ss.slice[1:], appendVal)
    } else {
        ss.slice = append(ss.slice, appendVal)
    }
}

func (ss *SliceString) ToSlice() (slice []string)  {
    slice = ss.slice
    return
}

func (ss *SliceString) SetMaxLen(maxLen int) {
    ss.maxLen = maxLen
}
