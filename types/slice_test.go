package types

import (
    "testing"
    "fmt"
)

func TestSliceString_Remove(t *testing.T) {
    ss := make(SliceString, 0)
    ss = append(ss, "assdd")
    ss = append(ss, "vvvv")
    (&ss).Remove("vvvv")

    fmt.Printf("ss:%+v\n", &ss)
}
