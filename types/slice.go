package types

import "sync"

type SliceString []string

var SliceStrMutex sync.Mutex

func NewSliceString() (ss SliceString) {
    ss = make(SliceString, 0)
    return
}

func NewSliceStringFromSlice(arr []string) (ss SliceString) {
    ss = arr
    return
}

func (ss *SliceString) Remove(rmVal string) (removedIdx int) {
    SliceStrMutex.Lock()
    defer SliceStrMutex.Unlock()
    removedIdx = -1
    ssObj := *ss
    for i, val := range ssObj {
        if val == rmVal {
            removedIdx = i
            break
        }
    }

    if removedIdx > -1 {
        ssObj = append(ssObj[:removedIdx], ssObj[removedIdx+1:]...)
    }
    *ss = ssObj
    return
}

func (ss *SliceString) Append(appendVal string)  {
    *ss = append(*ss, appendVal)
}

func (ss *SliceString) ToSlice() (slice []string)  {
    slice = []string(*ss)
    return
}
