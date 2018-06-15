package types

type SliceString []string

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

func (ss *SliceString) Append(appendVal string) (err error)  {

    //*ss = append()

}
