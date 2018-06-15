package types

type SliceString []string

func NewSliceString() (ss SliceString) {
    ss = make(SliceString, 0)
    return
}

func (ss *SliceString) Remove(rmVal string) {
    index := -1
    ssObj := *ss
    for i, val := range ssObj {
        if val == rmVal {
            index = i
            break
        }
    }

    if index > -1 {
        ssObj = append(ssObj[:index], ssObj[index+1:]...)
    }

    *ss = ssObj
}

func (ss *SliceString) Append(appendVal string)  {
    *ss = append(*ss, appendVal)
}

func (ss *SliceString) ToSlice() (slice []string)  {
    slice = []string(*ss)
    return
}
