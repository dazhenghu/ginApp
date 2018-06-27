package identify

import "github.com/dchest/captcha"

func Verify(id string, digits []byte) bool  {
    return captcha.Verify(id, digits)
}

func VerifyString(id string, digits string) bool {
    return captcha.VerifyString(id, digits)
}
