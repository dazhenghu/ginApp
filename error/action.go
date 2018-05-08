package error

/**
action执行前错误
 */
type BeforeActionErr interface {
    BeforeActionErrDes() string
}

/**
action执行后错误
 */
type AfterActionErr interface {
    AfterActionErrDes() string
}