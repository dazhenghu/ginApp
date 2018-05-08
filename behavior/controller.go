package behavior

import "github.com/gin-gonic/gin"

type Controller interface {
    beforeAction(context *gin.Context)
    afterAction(context *gin.Context)
}

