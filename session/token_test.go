package session

import (
    "testing"
    "encoding/json"
    "fmt"
)


func TestGenerateSessionToken(t *testing.T) {
    str := "[\"a\", \"b\"]"
    var arr []string
    err := json.Unmarshal([]byte(str), &arr)
    fmt.Printf("err:%+v\n", err)
    fmt.Printf("arr:%+v\n", arr)

    strEncode, _ := json.Marshal([]string{"rr", "ddd"})
    fmt.Printf("strEncode:%+v\n", string(strEncode))
}
