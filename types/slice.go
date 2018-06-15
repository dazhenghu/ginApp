package types

import "sync"

type SliceString struct {
    slice []string
    sync.Mutex
}

func NewSliceString() (ss *SliceString) {
    ss = &SliceString{
        slice: make([]string, 0),
    }
    return
}

func NewSliceStringFromSlice(arr []string) (ss *SliceString) {
    ss = &SliceString{
        slice: arr,
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
    ss.slice = append(ss.slice, appendVal)
}

func (ss *SliceString) ToSlice() (slice []string)  {
    slice = ss.slice
    return
}
