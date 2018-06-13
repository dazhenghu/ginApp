package token

import (
    "testing"
    "fmt"
)

func TestNewToken(t *testing.T) {
    tokenObj := NewToken("")
    fmt.Printf("token:%s\n", tokenObj.GenerateToken())
}
